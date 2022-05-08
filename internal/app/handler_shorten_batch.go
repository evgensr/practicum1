package app

import (
	"encoding/json"
	"fmt"
	"github.com/evgensr/practicum1/internal/helper"
	"log"
	"net/http"
)

// HandlerShortenBatch получить по хешу ссылку
func (s *APIserver) HandlerShortenBatch() http.HandlerFunc {

	type LineRequest struct {
		CorrelationID string `json:"correlation_id"`
		OriginalUrl   string `json:"original_url"`
	}

	type LineResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortUrl      string `json:"short_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// id пользователя
		userId := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))

		// создаем перменную
		decoder := json.NewDecoder(r.Body)

		// объявляем переменную запроса
		var lineRequest []LineRequest
		var lineResponse []LineResponse

		// декодируем в структуру request
		err := decoder.Decode(&lineRequest)
		if err != nil {
			// логируем ошибку
			s.logger.Warning("HandlerSetURL: request not json ", err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		log.Println(lineRequest)

		for _, line := range lineRequest {
			// получаем short
			hash := helper.GetShort(line.OriginalUrl)
			// записываем в хранилище ключ значение
			s.store.Set(line.OriginalUrl, hash, userId)
			log.Println(line)

			lineResponse = append(lineResponse, LineResponse{
				CorrelationID: line.CorrelationID,
				ShortUrl:      s.config.BaseURL + hash,
			})
		}

		// заголов ответа json
		w.Header().Set("Content-Type", "application/json")

		// заголов ответа 201
		w.WriteHeader(http.StatusCreated)

		// пишем в http.ResponseWriter ответ json
		err = json.NewEncoder(w).Encode(lineResponse)
		if err != nil {
			// заголов ответа 500
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

}
