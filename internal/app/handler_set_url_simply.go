package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/evgensr/practicum1/internal/helper"
)

// HandlerSetURLSimply creating an entry for the url
func (s *APIserver) HandlerSetURLSimply() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		hash := helper.GetShort(string(b))

		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		// записываем в хранилище ключ значение
		if err := s.store.Set(Line{
			URL:   string(b),
			Short: hash,
			User:  userID,
		}); err != nil {
			// заголовок ответа 409
			w.WriteHeader(http.StatusConflict)
		} else {
			// заголовок ответа 201
			w.WriteHeader(http.StatusCreated)
		}

		w.Write([]byte(s.config.BaseURL + hash))

	}

}
