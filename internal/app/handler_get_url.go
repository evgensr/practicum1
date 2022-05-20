package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// HandlerGetURL получить по хешу ссылку
func (s *APIserver) HandlerGetURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// spew.Dump(vars)
		// spew.Dump(r.Header)

		if len(vars["hash"]) < 1 {
			s.logger.Info("HandlerGetURL: vars hash is small ", vars["hash"])
			w.WriteHeader(500)
			return
		}

		url, err := s.store.Get(vars["hash"])
		if err != nil {
			log.Println(err)
		}

		// s.store.Debug()

		if len(url.Short) > 0 && url.Status != 1 {
			s.logger.Info("HandlerGetURL: result ", vars["hash"])
			w.Header().Set("Location", url.Short)
			w.WriteHeader(307)
			// w.Write(nil)
		} else if url.Status == 1 {
			w.WriteHeader(http.StatusGone)
		} else {
			s.logger.Info("HandlerGetURL: not found ", vars["hash"])
			w.WriteHeader(400)
		}
	}

}
