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
		URL   string `json:"original_url"`
		Short string `json:"short_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var line []Line
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
