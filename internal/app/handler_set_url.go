package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/evgensr/practicum1/internal/helper"
)

// HandlerSetURL creating an entry for the url
func (s *APIserver) HandlerSetURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		// объявляем переменную запроса
		var request request

		// проверяем на валидность URL
		_, err := govalidator.ValidateStruct(decoder)
		if err != nil {
			s.logger.Warning("HandlerSetURL: validator fail   ", err.Error())
			return
		}

		// декодируем в структуру request
		err = decoder.Decode(&request)
		if err != nil {
			s.logger.Warning("HandlerSetURL: request not json ", err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		// получаем short
		hash := helper.GetShort(request.URL)

		// заголовок ответа json
		w.Header().Set("Content-Type", "application/json")

		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		// записываем в хранилище ключ значение
		if err := s.store.Set(Line{
			URL:   request.URL,
			User:  userID,
			Short: hash,
		}); err != nil {
			// заголовок ответа 409
			w.WriteHeader(http.StatusConflict)
		} else {
			// заголовок ответа 201
			w.WriteHeader(http.StatusCreated)
		}

		// создаем переменную из структуры response
		data := response{
			URL: s.config.BaseURL + hash,
		}

		// пишем в http.ResponseWriter ответ json
		json.NewEncoder(w).Encode(data)

	}

}
