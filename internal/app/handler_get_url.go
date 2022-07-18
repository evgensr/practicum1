package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerGetURL получить по хешу ссылку
func (s *APIserver) HandlerGetURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if len(vars["hash"]) < 1 {
			s.logger.Info("HandlerGetURL: vars hash is small ", vars["hash"])
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		url, err := s.store.Get(vars["hash"])
		if err != nil {
			log.Println(err)
		}

		if len(url.Short) > 0 && url.Status != 1 {
			s.logger.Info("HandlerGetURL: result ", vars["hash"])
			w.Header().Set("Location", url.URL)
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else if url.Status == 1 {
			w.WriteHeader(http.StatusGone)
		} else {
			s.logger.Info("HandlerGetURL: not found ", vars["hash"])
			w.WriteHeader(http.StatusBadRequest)
		}
	}

}
