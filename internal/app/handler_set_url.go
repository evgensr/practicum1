package app

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/evgensr/practicum1/internal/helper"
	"net/http"
)

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

		// получаем short
		hash := helper.GetShort(request.URL)

		// id пользователя
		userId := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		// записываем в хранилище ключ значение
		s.store.Set(request.URL, hash, userId)

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
