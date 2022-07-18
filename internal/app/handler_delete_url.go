package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HandlerDeleteURL получить по хешу ссылку
func (s *APIserver) HandlerDeleteURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("delete")
		// создаем переменную
		decoder := json.NewDecoder(r.Body)

		// объявляем переменную запроса
		var lineRequest []string
		// var line []Line
		// id пользователя
		userID := fmt.Sprintf("%v", r.Context().Value(ctxKeyUser))

		// декодируем в структуру request
		err := decoder.Decode(&lineRequest)

		log.Println("lineRequest ", lineRequest)

		if err != nil {
			log.Println("lineRequest err", err)
			return
		}

		for _, row := range lineRequest {
			go s.store.Delete([]Line{
				{
					Short: row,
					User:  userID,
				},
			})
		}

		w.WriteHeader(http.StatusAccepted)

	}

}
