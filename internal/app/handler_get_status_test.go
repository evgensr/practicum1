package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandlerGetStatus(t *testing.T) {

	conf := NewConfig()
	// sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf)

	t.Run("stats", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/api/internal/stats", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		s.HandlerGetStatus().ServeHTTP(rr, req)
		router.HandleFunc("/api/internal/stats", s.HandlerGetStatus())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK)
	})

}
