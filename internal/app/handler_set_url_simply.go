package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/evgensr/practicum1/internal/helper"
)

// HandlerSetURLSimply - создаем запись для url
func (s *APIserver) HandlerSetURLSimply() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		log.Println(string(b))

		// получаем short
		hash := helper.GetShort(string(b))

		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))
		// записываем в хранилище ключ значение
		if err := s.store.Set(Line{
			URL:   string(b),
			Short: hash,
			User:  userID,
		}); err != nil {
			// заголов ответа 409
			w.WriteHeader(http.StatusConflict)
		} else {
			// заголов ответа 201
			w.WriteHeader(http.StatusCreated)
		}

		w.Write([]byte(s.config.BaseURL + hash))

	}

}
