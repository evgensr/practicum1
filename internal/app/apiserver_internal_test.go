package app

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_Handler(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)
	rec := httptest.NewRecorder()

	values := map[string]string{"url": "https://habr8234.ru"}
	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"response\":\"61f88ac1dd3c6bf3d9f6e4e15e454288\"}\n")

}
