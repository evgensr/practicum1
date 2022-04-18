package app

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/evgensr/practicum1/internal/helper"
	"github.com/evgensr/practicum1/internal/store/file"
	"github.com/evgensr/practicum1/internal/store/memory"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// store  memory.RWMap
	// store file.RWMap
	store Storage
}

// New ...
func New(config *Config) *APIserver {
	var store Storage

	// TODO: уловие в зависимости от конфига
	param := config.FileStoragePath
	if len(param) > 1 {
		store = file.New(param)
	} else {
		store = memory.New(param)
	}

	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		// store:  memory.New(),
		store: store,
	}
}

// Start ...
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("SERVER_ADDRESS ", s.config.ServerAddress)
	s.logger.Info("BASE_URL ", s.config.BaseURL)
	s.logger.Info("FILE_STORAGE_PATH ", s.config.FileStoragePath)

	s.logger.Info("Staring api server")
	return http.ListenAndServe(s.config.ServerAddress, s.router)
}

func (s *APIserver) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil

}

func (s *APIserver) configureRouter() {

	s.router.HandleFunc("/{hash}", s.HandlerGetURL)
	s.router.HandleFunc("/api/shorten", s.HandlerSetURL).Methods("POST")

}

func (s *APIserver) configureStore() error {
	return nil
}

// HandlerGetURL получить по хешу ссылку
func (s *APIserver) HandlerGetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if len(vars["hash"]) != 32 {
		s.logger.Info("HandlerGetURL: vars hash is small ", vars["hash"])
		w.WriteHeader(500)
		return
	}

	url := s.store.Get(vars["hash"])

	// s.store.Debug()

	if len(url) > 0 {
		s.logger.Info("HandlerGetURL: result ", vars["hash"])
		w.Header().Set("Location", url)
		w.WriteHeader(307)
		// w.Write(nil)
	} else {
		s.logger.Info("HandlerGetURL: not found ", vars["hash"])
		w.WriteHeader(400)
	}

}

// HandlerSetURL - создаем запись для url
func (s *APIserver) HandlerSetURL(w http.ResponseWriter, r *http.Request) {

	// создаем перменную
	decoder := json.NewDecoder(r.Body)

	// объявляем переменную запроса
	var request request

	// проверяем на валидность URL
	_, err := govalidator.ValidateStruct(decoder)
	if err != nil {
		// логируем ошибку
		s.logger.Warning("HandlerSetURL: validator fail   ", err.Error())
		return
	}

	// декодируем в структуру request
	err = decoder.Decode(&request)
	if err != nil {
		// логируем ошибку
		s.logger.Warning("HandlerSetURL: request not json ", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// spew.Dump(request.URL)

	// получаем хеш url для записи
	hash := helper.GetHash(request.URL)

	// записываем в хранилище ключ значение
	s.store.Set(hash, request.URL)

	// создаем переменную из структуры response
	data := response{
		URL: s.config.BaseURL + hash,
	}
	// записываем в лог ответ
	s.logger.Info("HandlerSetURL: response ", data)

	// заголов ответа json
	w.Header().Set("Content-Type", "application/json")

	// заголов ответа 201
	w.WriteHeader(http.StatusCreated)

	// пишем в http.ResponseWriter ответ json
	json.NewEncoder(w).Encode(data)

}
