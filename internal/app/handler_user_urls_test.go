package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandlerUserUrls(t *testing.T) {

	conf := NewConfig()
	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf, sessionStore)

	t.Run("create", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/api/user/urls", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		s.HandlerUserUrls().ServeHTTP(rr, req)
		router.HandleFunc("/api/user/urls", s.HandlerUserUrls())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusNoContent)
	})

}
