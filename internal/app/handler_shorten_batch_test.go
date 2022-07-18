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

func TestAPIserver_HandlerShortenBatch(t *testing.T) {

	type LineRequest struct {
		CorrelationID string `json:"correlation_id"`
		OriginalURL   string `json:"original_url"`
	}

	conf := NewConfig()
	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf, sessionStore)

	t.Run("create", func(t *testing.T) {

		URL = "https://habr.com"
		lineRequest := []LineRequest{
			{
				CorrelationID: "123",
				OriginalURL:   URL,
			},
		}
		jsonValue, _ := json.Marshal(lineRequest)
		req, err := http.NewRequest(http.MethodPost, "/api/shorten/batch", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		s.HandlerShortenBatch().ServeHTTP(rr, req)
		router.HandleFunc("/api/shorten/batch", s.HandlerGetURL())
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusCreated)
	})

}
