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
		originalURL   string `json:"original_url"`
	}

	type LineResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))

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
			hash := helper.GetShort(line.originalURL)
			// записываем в хранилище ключ значение
			s.store.Set(line.originalURL, hash, userID)
			log.Println(line)

			lineResponse = append(lineResponse, LineResponse{
				CorrelationID: line.CorrelationID,
				ShortURL:      s.config.BaseURL + hash,
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
