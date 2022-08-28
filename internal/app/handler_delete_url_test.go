package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evgensr/practicum1/internal/helper"
	"github.com/stretchr/testify/assert"
)

var (
	URL   string
	Short string
)

func TestAPIserver_HandlerDeleteURL(t *testing.T) {

	URL = "https://habr.ru"
	Short = helper.GetShort(URL)

	conf := NewConfig()
	// sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := New(&conf)

	type response struct {
		URL string `json:"result" valid:"url"`
	}

	data := response{
		URL: s.config.BaseURL + Short,
	}
	// json.NewEncoder(w).Encode(data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Log(err)
		return
	}

	record := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(map[string]string{"url": URL})
	req, _ := http.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonValue))
	s.HandlerSetURL().ServeHTTP(record, req)
	assert.Equal(t, record.Body.String(), string(b)+"\n")

	record = httptest.NewRecorder()
	jsonValue, _ = json.Marshal([]string{Short})
	req, _ = http.NewRequest(http.MethodDelete, "/api/user/urls", bytes.NewBuffer(jsonValue))
	s.HandlerDeleteURL().ServeHTTP(record, req)
	t.Log("code", record.Code)
	assert.Equal(t, record.Code, http.StatusAccepted)
}
