package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/evgensr/practicum1/internal/store"
)

// HandlerUserUrls
func (s *APIserver) HandlerUserUrls() http.HandlerFunc {

	type Line struct {
		Short string `json:"short_url"`
		URL   string `json:"original_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var line []Line
		var bLine []store.Line
		// заголовок ответа json
		w.Header().Set("Content-Type", "application/json")
		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		bLine = s.store.GetByUser(userID)
		if len(bLine) > 0 {
			for _, url := range bLine {
				line = append(line, Line{
					URL:   url.URL,
					Short: s.config.BaseURL + url.Short,
				})
			}

			log.Println(line)
			// пишем в http.ResponseWriter ответ json
			json.NewEncoder(w).Encode(line)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}

	}

}
