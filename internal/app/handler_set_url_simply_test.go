package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandlerSetURLSimply(t *testing.T) {

	conf := NewConfig()
	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf, sessionStore)

	t.Run("create", func(t *testing.T) {
		URL = "https://habr.com"
		jsonValue, _ := json.Marshal(map[string]string{"url": URL})
		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		s.HandlerSetURLSimply().ServeHTTP(rr, req)
		router.HandleFunc("/", s.HandlerGetURL())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusCreated)
	})

}
