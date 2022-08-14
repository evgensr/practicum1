package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIserver_GzipHandleEncode(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	authenticateHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	res := httptest.NewRecorder()

	http.SetCookie(res, &http.Cookie{Name: sessionName, Value: "expected"})

	authenticateHandler(res, req)

	tim := s.GzipHandleEncode(authenticateHandler)
	tim.ServeHTTP(res, req)

}

func TestAPIserver_GzipHandleDecode(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	authenticateHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	res := httptest.NewRecorder()

	http.SetCookie(res, &http.Cookie{Name: sessionName, Value: "expected"})

	authenticateHandler(res, req)

	tim := s.GzipHandleDecode(authenticateHandler)
	tim.ServeHTTP(res, req)

}
