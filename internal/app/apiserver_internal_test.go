package app

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_Handler(t *testing.T) {

	conf := NewConfig()
	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))

	s := New(&conf, sessionStore)

	rec := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(map[string]string{"url": "https://habr8234.ru"})
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"result\":\"61f88ac1dd3c6bf3d9f6e4e15e454288\"}\n")

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
		url          string
		method       string
	}{
		{
			name: "valid",
			payload: map[string]string{
				"url": "http://yandex.ru",
			},
			expectedCode: http.StatusCreated,
			url:          "/api/shorten",
			method:       http.MethodPost,
		},
		{
			name: "valid StatusConflict",
			payload: map[string]string{
				"url": "http://yandex.ru",
			},
			expectedCode: http.StatusConflict,
			url:          "/api/shorten",
			method:       http.MethodPost,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, tc.url, b)
			s.HandlerSetURL().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
