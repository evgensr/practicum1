package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log"
	"net/http"
)

const (
	port     = "8080"
	host     = "localhost"
	protocol = "http"
)

var mapURL = make(map[string]string)

// handlerURL — обработчик запроса.
func handlerURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	// если методом POST
	case "POST":
		log.Println("post request")

		// читаем Body
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		log.Println(string(b))
		address := string(b)
		hash := GetHash(address)
		log.Println(hash)
		mapURL[hash] = address
		log.Printf("Address = %s\n", address)
		w.WriteHeader(201)
		w.Write([]byte(protocol + "://" + host + ":" + port + "/" + hash))

	case "GET":
		urlHash := r.RequestURI[1:]
		log.Println(urlHash)

		val, exists := mapURL[urlHash]
		log.Println(mapURL)
		if exists {
			log.Println(val)
			w.Header().Set("Location", val)
			w.WriteHeader(307)
			w.Write(nil)
		} else {
			delete(mapURL, urlHash)
			w.WriteHeader(400)
		}
	}

}

// HandlerPOST — обработчик запроса.
func HandlerPOST(w http.ResponseWriter, r *http.Request) {
	log.Println("post request")
	// читаем Body
	b, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(string(b))
	address := string(b)

	hash := GetHash(address)
	log.Println(hash)
	mapURL[hash] = address
	log.Printf("Address = %s\n", address)
	w.WriteHeader(201)
	w.Write([]byte(protocol + "://" + host + ":" + port + "/" + hash))

}

// HandlerGET — обработчик запроса.
func HandlerGET(w http.ResponseWriter, r *http.Request) {

	urlHash := chi.URLParam(r, "hash")
	log.Println(urlHash)
	// len := len(urlHash)

	if len(urlHash) != 32 {
		w.WriteHeader(500)
		return
	}

	// Вопрос, как еще можно проверить сущесовование map с таким ключом, без создания?
	val, exists := mapURL[urlHash]
	// log.Println(mapURL)
	if exists {
		log.Println(val)
		w.Header().Set("Location", val)
		w.WriteHeader(307)
		w.Write(nil)
	} else {
		delete(mapURL, urlHash)
		w.WriteHeader(400)
	}

}

func GetHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

	log.Println("start server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{hash}", HandlerGET)
	r.Post("/", HandlerPOST)

	//// маршрутизация запросов обработчику
	//http.HandleFunc("/", handlerURL)

	// запуск сервера с адресом localhost, порт port
	// err := http.ListenAndServe(":"+port, nil)
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatal(err)
	}

}
