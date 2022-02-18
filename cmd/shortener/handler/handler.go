package handler

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

var mapURL = make(map[string]string)

const (
	Port     = "8080"
	Host     = "localhost"
	Protocol = "http"
)

func GetHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
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

// HandlerPOST — обработчик запроса.
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	message := make(chan string)
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
	// конкуретная запись
	// рецепт взять отсюда https://youtu.be/wHQBMDInWEg?t=354
	go func() {
		message <- address
	}()
	// mapURL[hash] = address
	mapURL[hash] = <-message
	log.Printf("Address = %s\n", address)
	w.WriteHeader(201)
	w.Write([]byte(Protocol + "://" + Host + ":" + Port + "/" + hash))

}
