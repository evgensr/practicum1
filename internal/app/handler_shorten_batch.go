package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evgensr/practicum1/internal/helper"
)

// HandlerShortenBatch get a hash link
func (s *APIserver) HandlerShortenBatch() http.HandlerFunc {

	type LineRequest struct {
		CorrelationID string `json:"correlation_id"`
		OriginalURL   string `json:"original_url"`
	}

	type LineResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))

		decoder := json.NewDecoder(r.Body)

		// объявляем переменную запроса
		var lineRequest []LineRequest
		var lineResponse []LineResponse

		// декодируем в структуру request
		err := decoder.Decode(&lineRequest)
		if err != nil {
			s.logger.Warning("HandlerSetURL: request not json ", err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		for _, line := range lineRequest {
			// получаем short
			hash := helper.GetShort(line.OriginalURL)
			// записываем в хранилище ключ значение
			s.store.Set(Line{
				URL:           line.OriginalURL,
				Short:         hash,
				User:          userID,
				CorrelationID: line.CorrelationID,
			})

			lineResponse = append(lineResponse, LineResponse{
				CorrelationID: line.CorrelationID,
				ShortURL:      s.config.BaseURL + hash,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		// пишем в http.ResponseWriter ответ json
		err = json.NewEncoder(w).Encode(lineResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

}
