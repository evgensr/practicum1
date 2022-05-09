package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HandlerUserUrls
func (s *APIserver) HandlerUserUrls() http.HandlerFunc {

	type Line struct {
		Short string `json:"short_url"`
		URL   string `json:"original_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var line []Line
		// заголов ответа json
		w.Header().Set("Content-Type", "application/json")
		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		urls := s.store.GetByUser(userID)
		if len(urls) > 0 {
			for _, url := range urls {
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
