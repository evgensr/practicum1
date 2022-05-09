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

	// conf := NewConfig()
	// s := New(&conf)
	rec := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(map[string]string{"url": "https://habr8234.ru"})
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"result\":\"61f88ac1dd3c6bf3d9f6e4e15e454288\"}\n")

}
