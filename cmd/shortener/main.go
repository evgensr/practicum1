package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
)

const (
	port = "8080"
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
		w.Write([]byte("http://localhost:" + port + "/" + hash))

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

func GetHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

	log.Println("start server")

	// маршрутизация запросов обработчику
	http.HandleFunc("/", handlerURL)

	// запуск сервера с адресом localhost, порт port
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
