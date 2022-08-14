package app

import (
	"database/sql"
	"log"
	"net/http"
)

// HandlerPing проверка подключения к базе данных
func (s *APIserver) HandlerPing() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("DatabaseDSN: ", s.config.DatabaseDSN)
		db, err := sql.Open("postgres", s.config.DatabaseDSN)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := db.Ping(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}

}
