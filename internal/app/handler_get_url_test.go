package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evgensr/practicum1/internal/helper"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandlerGetURL(t *testing.T) {

	URL = "https://habr.ru"
	Short = helper.GetShort(URL)

	conf := NewConfig()
	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf, sessionStore)
	record := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(map[string]string{"url": URL})
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(record, req)
	assert.Equal(t, record.Body.String(), "{\"result\":\""+Short+"\"}\n")

	t.Run("redirect 307", func(t *testing.T) {
		path := fmt.Sprintf("/%s", Short)
		req, err := http.NewRequest(http.MethodGet, path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/{hash}", s.HandlerGetURL())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusTemporaryRedirect)
	})

	t.Run("error 500", func(t *testing.T) {
		record = httptest.NewRecorder()
		reqGetURL, _ := http.NewRequest(http.MethodGet, "/1", nil)
		s.HandlerGetURL().ServeHTTP(record, reqGetURL)
		t.Log("code", record.Code)
		assert.Equal(t, record.Code, http.StatusInternalServerError)
	})

	t.Run("status bad request", func(t *testing.T) {
		Short = helper.GetShort("test")
		path := fmt.Sprintf("/%s", Short)
		req, err := http.NewRequest(http.MethodGet, path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/{hash}", s.HandlerGetURL())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})

}
