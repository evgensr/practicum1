package app

import (
	"github.com/evgensr/practicum1/internal/store/file"
	"github.com/evgensr/practicum1/internal/store/memory"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
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

	s.router.HandleFunc("/{hash}", s.HandlerGetURL())
	s.router.HandleFunc("/api/shorten", s.HandlerSetURL()).Methods("POST")
	s.router.Use(s.GzipHandle)

}

func (s *APIserver) configureStore() error {
	return nil
}
