package app

import (
	"encoding/json"
	"net/http"

	"github.com/evgensr/practicum1/internal/store"
)

// HandlerGetStatus getting the number of users and links
func (s *APIserver) HandlerGetStatus() http.HandlerFunc {

	type Response struct {
		Urls  store.Urls  `json:"urls"`
		Users store.Users `json:"users"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var res Response

		urls, users, err := s.store.GetStats()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res = Response{
			Users: users,
			Urls:  urls,
		}

		json.NewEncoder(w).Encode(res)

	}

}
