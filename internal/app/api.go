// Package app for launching a web server
package app

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/evgensr/practicum1/internal/helper"
	"github.com/evgensr/practicum1/internal/store/file"
	"github.com/evgensr/practicum1/internal/store/memory"
	"github.com/evgensr/practicum1/internal/store/pg"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	shutdownTimeout = 5 * time.Second
)

// APIserver required components for the server
type APIserver struct {
	config       *Config        //server config
	logger       *logrus.Logger //logrus is a structured logger for Go
	router       *mux.Router    //gorilla/mux
	store        Storage        //interface for working with storage
	sessionStore sessions.Store //interface for working with sessions
	closer       *Closer
}

// New Creates a new app.
func New(config *Config) *APIserver {
	var store Storage
	logger := logrus.New()

	c := &Closer{}

	// sessionStore = sessions.NewCookieStore([]byte(config.SessionKey))
	config.BaseURL = helper.AddSlash(config.BaseURL)
	param := config.FileStoragePath
	if len(config.DatabaseDSN) > 1 {
		store = pg.New(config.DatabaseDSN)
		logger.Info("store pg")
	} else if len(param) > 1 {
		store = file.New(param)
		logger.Info("store file")
	} else {
		store = memory.New(param)

		logger.Info("store memory")
	}

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	return &APIserver{
		config:       config,
		logger:       logger,
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
		closer:       c,
	}
}

// Start web server
func (s *APIserver) Start(ctx context.Context) error {

	// c := &Closer{}

	// var mux = http.NewServeMux()
	var srv = &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.router,
	}

	s.closer.Add(srv.Shutdown)
	s.closer.Add(s.store.Shutdown)

	if err := s.configureLogger(); err != nil {
		return err
	}

	if len(s.config.DatabaseDSN) > 1 {

		if err := pg.RunMigrations(s.config.DatabaseDSN, "file://internal/store/pg/migrations/"); err != nil {
			log.Fatal("create table ", err)
		}
	}

	s.configureRouter()

	s.logger.Info("SERVER_ADDRESS ", s.config.ServerAddress)
	s.logger.Info("BASE_URL ", s.config.BaseURL)
	s.logger.Info("FILE_STORAGE_PATH ", s.config.FileStoragePath)
	s.logger.Info("ENABLE_HTTPS ", s.config.EnableHTTPS)

	if bytes.Contains([]byte(s.config.BaseURL), []byte("http://")) && s.config.EnableHTTPS {
		s.logger.Warning("warn: you need use HTTPS in base_url")
	}

	go func() {
		if s.config.EnableHTTPS {
			if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
				s.logger.Info(err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil {
				s.logger.Info(err)
			}
		}

	}()

	s.logger.Infof("Staring api server")
	<-ctx.Done()

	s.logger.Infof("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.closer.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil

}

//configureLogger setup logrus
func (s *APIserver) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil

}

//configureRouter setup router
func (s *APIserver) configureRouter() {

	s.router.Use(s.authenticateUser)
	s.router.HandleFunc("/ping", s.HandlerPing())
	s.router.HandleFunc("/api/user/urls", s.HandlerUserUrls()).Methods(http.MethodGet)
	s.router.HandleFunc("/{hash}", s.HandlerGetURL()).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.HandlerSetURLSimply()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/shorten", s.HandlerSetURL()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/shorten/batch", s.HandlerShortenBatch()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/urls", s.HandlerDeleteURL()).Methods(http.MethodDelete)
	s.router.Use(s.GzipHandleEncode)
	s.router.Use(s.GzipHandleDecode)
	// s.router.Use(s.Log)

}

//authenticateUser assigns a cookie to a new client
//if there is already a cookie, decodes it and writes it to ctxKeyUser
func (s *APIserver) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var id string

		s.logger.Info("SessionKey:  ", s.config.SessionKey)

		c, err := r.Cookie(sessionName)
		if err != nil {
			expiration := time.Now().Add(365 * 24 * time.Hour)
			id = helper.GeneratorUUID()
			encryptedCookie, err := helper.Encrypted([]byte(id), s.config.SessionKey)
			if err != nil {
				s.logger.Warning("error encrypted cookie ", err)
				return
			}
			cookie := http.Cookie{Name: sessionName, Value: hex.EncodeToString(encryptedCookie), Expires: expiration}
			http.SetCookie(w, &cookie)
		} else {
			// fmt.Println("Cookie ", c.Value)
			s.logger.Info("Cookie ", c.Value)
			decoded, err := hex.DecodeString(c.Value)
			if err != nil {
				s.logger.Warning("error decode string Cookie ", c.Value)
				return
			}
			decryptedCookie, err := helper.Decrypted(decoded, s.config.SessionKey)
			if err != nil {
				// не смогли декодировать, устанавливаем новую куку и юзера
				expiration := time.Now().Add(365 * 24 * time.Hour)
				id = helper.GeneratorUUID()
				encryptedCookie, err := helper.Encrypted([]byte(id), s.config.SessionKey)
				if err != nil {
					s.logger.Warning("error encrypted cookie ", err)
					return
				}
				cookie := http.Cookie{Name: sessionName, Value: hex.EncodeToString(encryptedCookie), Expires: expiration}
				http.SetCookie(w, &cookie)

			}

			id = string(decryptedCookie)
		}

		s.logger.Info("user id: ", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, id)))

	})

}

//error server error handling
func (s *APIserver) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

//respond server response wrapper
func (s *APIserver) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

//CreateTable creates a short table if it does not exist
//func (s *APIserver) CreateTable() error {
//	log.Println("config Database: ", s.config.DatabaseDSN)
//	db, err := sql.Open("postgres", s.config.DatabaseDSN)
//	if err != nil {
//		log.Println("create table func ", err)
//		return err
//	}
//	if err := db.Ping(); err != nil {
//		log.Println("ping err ", err)
//		return err
//	}
//
//	if _, err := db.Exec("CREATE TABLE  IF NOT EXISTS short" +
//		"(id serial primary key," +
//		"original_url varchar(4096) not null," +
//		"short_url varchar(32) UNIQUE not null," +
//		"user_id varchar(36) not null," +
//		"correlation_id varchar(36) null," +
//		"status smallint not null DEFAULT 0);"); err != nil {
//		return errors.New("error sql ")
//	}
//
//	return nil
//
//}
