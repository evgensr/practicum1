package app

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/evgensr/practicum1/internal/helper"
	"github.com/evgensr/practicum1/internal/store/file"
	"github.com/evgensr/practicum1/internal/store/memory"
	"github.com/evgensr/practicum1/internal/store/pg"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

// APIserver ...
type APIserver struct {
	config       *Config
	logger       *logrus.Logger
	router       *mux.Router
	store        Storage
	sessionStore sessions.Store
}

// New ...
func New(config *Config, sessionStore sessions.Store) *APIserver {
	var store Storage
	// sessionStore = sessions.NewCookieStore([]byte(config.SessionKey))
	config.BaseURL = helper.AddSlash(config.BaseURL)
	param := config.FileStoragePath
	if len(config.DatabaseDSN) > 1 {
		store = pg.New(config.DatabaseDSN)
		log.Println("store pg")
	} else if len(param) > 1 {
		store = file.New(param)
		log.Println("store file")
	} else {
		store = memory.New(param)
		log.Println("store memory")
	}

	return &APIserver{
		config:       config,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}
}

// Start ...
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	if len(s.config.DatabaseDSN) > 1 {
		if err := s.CreateTable(); err != nil {
			log.Fatal("create table ", err)
		}
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		log.Println("configureStore ", err)
		return err
	}

	//args := os.Args
	//fmt.Printf("All arguments: %v\n", args)

	s.logger.Info("SERVER_ADDRESS ", s.config.ServerAddress)
	s.logger.Info("BASE_URL ", s.config.BaseURL)
	s.logger.Info("FILE_STORAGE_PATH ", s.config.FileStoragePath)

	s.logger.Info("Staring api server")
	return http.ListenAndServe(s.config.ServerAddress, s.router)
}

func (s *APIserver) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil

}

func (s *APIserver) configureRouter() {

	s.router.Use(s.authenticateUser)
	s.router.HandleFunc("/ping", s.HandlerPing())
	s.router.HandleFunc("/api/user/urls", s.HandlerUserUrls()).Methods(http.MethodPost)
	s.router.HandleFunc("/{hash}", s.HandlerGetURL()).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.HandlerSetURLSimply()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/shorten", s.HandlerSetURL()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/shorten/batch", s.HandlerShortenBatch()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/urls", s.HandlerDeleteURL()).Methods(http.MethodDelete)
	s.router.Use(s.Gzip)
	// s.router.Use(s.Log)

}

func (s *APIserver) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var id string
		// log.Println("Header: ", r.Header)
		// log.Println("Body: ", r.Body)

		log.Println("SessionKey:  ", s.config.SessionKey)

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
			fmt.Println("Cookie ", c.Value)
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

		log.Println("user id: ", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, id)))

	})

}

func (s *APIserver) configureStore() error {
	return nil
}

func (s *APIserver) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

func (s *APIserver) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *APIserver) CreateTable() error {
	log.Println("config Database: ", s.config.DatabaseDSN)
	db, err := sql.Open("postgres", s.config.DatabaseDSN)
	if err != nil {
		log.Println("create table func ", err)
		return err
	}
	if err := db.Ping(); err != nil {
		log.Println("ping err ", err)
		return err
	}

	if _, err := db.Exec("CREATE TABLE  IF NOT EXISTS short" +
		"(id serial primary key," +
		"original_url varchar(4096) not null," +
		"short_url varchar(32) UNIQUE not null," +
		"user_id varchar(36) not null," +
		"correlation_id varchar(36) null," +
		"status smallint not null DEFAULT 0);"); err != nil {
		return errors.New("error sql ")
	}

	return nil

}
