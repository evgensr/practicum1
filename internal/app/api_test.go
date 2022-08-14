package app

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os/signal"
	"syscall"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAPIserver_configureRouter(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	type fields struct {
		config       *Config
		logger       *logrus.Logger
		router       *mux.Router
		store        Storage
		sessionStore sessions.Store
		closer       *Closer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s.configureRouter()
		})
	}
}

func TestAPIserver_configureLogger(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	type fields struct {
		config       *Config
		logger       *logrus.Logger
		router       *mux.Router
		store        Storage
		sessionStore sessions.Store
		closer       *Closer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.NoError(t, s.configureLogger())
		})
	}
}

func TestAPIserver_Start(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	type fields struct {
		config       *Config
		logger       *logrus.Logger
		router       *mux.Router
		store        Storage
		sessionStore sessions.Store
		closer       *Closer
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				stop()
			}()
			assert.NoError(t, s.Start(ctx))

		})
	}

}

func AddContextWithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func TestAPIserver_authenticateUser(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	//nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//
	//})

	// token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODcwNTIzMTcsImV4cCI6MTU4NzA1MjkxNywic2Vzc2lvbi1kYXRhIjoiVGVzdC5NY1Rlc3RGYWNlQG1haWwuY29tIn0.f0oM4fSH_b1Xi5zEF0VK-t5uhpVidk5HY1O0EGR4SQQ"
	req, err := http.NewRequest("GET", "/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	// req.Header.Set("Authorization", token)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Response.StatusCode == 403 {
			t.Fatalf("Response should be 200 for a valid JWT Token")
		}
	})

	rw := httptest.NewRecorder()
	handler := s.authenticateUser(testHandler)
	// 	handler.ServeHTTP(rw, req)
	log.Println(handler)
	log.Println(rw)
	log.Println(req)

}

func TestTimer(t *testing.T) {

	conf := NewConfig()
	s := New(&conf)

	authenticateHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	req := httptest.NewRequest(http.MethodGet, "/1", nil)

	res := httptest.NewRecorder()

	http.SetCookie(res, &http.Cookie{Name: sessionName, Value: "expected"})

	authenticateHandler(res, req)

	tim := s.authenticateUser(authenticateHandler)
	tim.ServeHTTP(res, req)
}
