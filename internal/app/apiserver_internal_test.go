package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_Handler(t *testing.T) {

	conf := NewConfig()
	// sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))

	s := New(&conf)

	rec := httptest.NewRecorder()

	type response struct {
		URL string `json:"result" valid:"url"`
	}

	// var request request

	data := response{
		URL: s.config.BaseURL + "61f88ac1dd3c6bf3d9f6e4e15e454288",
	}
	// json.NewEncoder(w).Encode(data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Log(err)
		return
	}

	jsonValue, _ := json.Marshal(map[string]string{"url": "https://habr8234.ru"})
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), string(b)+"\n")

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
