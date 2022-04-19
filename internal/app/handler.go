package app

import (
	"compress/gzip"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/evgensr/practicum1/internal/helper"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// HandlerGetURL получить по хешу ссылку
func (s *APIserver) HandlerGetURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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

}

// HandlerHello ...
func (s *APIserver) HandlerHello(w http.ResponseWriter, r *http.Request) {
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
func (s *APIserver) HandlerSetURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

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

}

// GzipHandle gzip-сжатие ответа
func (s *APIserver) GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			s.logger.Info("Not support gzip")
			next.ServeHTTP(w, r)
			return
		}

		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		s.logger.Info("Support gzip")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
