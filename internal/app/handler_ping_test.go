package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	t.Run("valid connect", func(t *testing.T) {

		dsn := os.Getenv("DATABASE_DSN")
		if dsn == "" {
			dsn = "postgres://postgres:postgres@localhost:5432/restapi?sslmode=disable"
			err := os.Setenv("DATABASE_DSN", dsn)
			assert.NoError(t, err)
			defer func() {
				err := os.Unsetenv("DATABASE_DSN")
				assert.NoError(t, err)
			}()
		}
		t.Log(os.Getenv("DATABASE_DSN"))

		conf := NewConfig()
		sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
		conf.DatabaseDSN = dsn
		s := New(&conf, sessionStore)
		rec := httptest.NewRecorder()
		b := &bytes.Buffer{}
		req, _ := http.NewRequest(http.MethodGet, "/ping", b)
		s.HandlerPing().ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("connection refused", func(t *testing.T) {

		dsn := "postgres://google:google@127.0.0.1:15432/restapi?sslmode=disable"
		conf := NewConfig()
		sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))
		conf.DatabaseDSN = dsn
		s := New(&conf, sessionStore)
		rec := httptest.NewRecorder()
		b := &bytes.Buffer{}
		req, _ := http.NewRequest(http.MethodGet, "/ping", b)
		s.HandlerPing().ServeHTTP(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
