package app

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/evgensr/practicum1/internal/helper"
	"github.com/evgensr/practicum1/internal/store/file"
	"github.com/evgensr/practicum1/internal/store/memory"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"time"
)

// APIserver ...
type APIserver struct {
	config       *Config
	logger       *logrus.Logger
	router       *mux.Router
	store        Storage
	sessionStore sessions.Store
	user         []User
}

// New ...
func New(config *Config, sessionStore sessions.Store) *APIserver {
	var store Storage
	// sessionStore = sessions.NewCookieStore([]byte(config.SessionKey))

	param := config.FileStoragePath
	if len(param) > 1 {
		store = file.New(param)
	} else {
		store = memory.New(param)
	}

	return &APIserver{
		config:       config,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
		user:         []User{},
	}
}

// Start ...
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	args := os.Args
	fmt.Printf("All arguments: %v\n", args)

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
	s.router.HandleFunc("/{hash}", s.HandlerGetURL())

	s.router.HandleFunc("/api/shorten", s.HandlerSetURL()).Methods("POST")
	s.router.Use(s.GzipHandle)

}

func (s *APIserver) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := ""

		c, err := r.Cookie(sessionName)
		if err != nil {
			expiration := time.Now().Add(365 * 24 * time.Hour)
			id = helper.GeneratorUuid()
			encryptedCookie, err := helper.Encrypted([]byte(id), "123")
			if err != nil {
				return
			}
			cookie := http.Cookie{Name: sessionName, Value: hex.EncodeToString(encryptedCookie), Expires: expiration}
			http.SetCookie(w, &cookie)
		} else {
			fmt.Println("Cookie ", c.Value)
			decoded, err := hex.DecodeString(c.Value)
			decryptedCookie, err := helper.Decrypted(decoded, "123")
			if err != nil {
				return
			}

			id = string(decryptedCookie)
		}

		log.Println(id)

		next.ServeHTTP(w, r)

		// а можно упростить все на mux
		// инициализация сессии
		//session, err := s.sessionStore.Get(r, sessionName)
		//if err != nil {
		//	s.error(w, r, http.StatusInternalServerError, err)
		//	return
		//}
		//
		//id, ok := session.Values["user_id"]
		//if !ok {
		//	session.Values["user_id"] = helper.GeneratorUuid()
		//	id = session.Values["user_id"]
		//	if err := s.sessionStore.Save(r, w, session); err != nil {
		//		s.error(w, r, http.StatusInternalServerError, err)
		//		return
		//	}
		//}
		//
		//log.Println(id)
		//
		//next.ServeHTTP(w, r)

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
