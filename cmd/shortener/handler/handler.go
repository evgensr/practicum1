package handler

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	Port     = "8080"
	Host     = "localhost"
	Protocol = "http"
)

// https://riptutorial.com/go/example/3423/concurrent-access-of-maps
// https://habr.com/ru/post/359078/

type RWMap struct {
	sync.RWMutex
	m map[string]string
}

func NewRWMap() *RWMap {
	return &RWMap{
		m: make(map[string]string),
	}
}

// Get is a wrapper for getting the value from the underlying map
func (c RWMap) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	return c.m[key]
}

// Set is a wrapper for setting the value of a key in the underlying map
func (c RWMap) Set(key string, val string) {
	c.Lock()
	defer c.Unlock()
	c.m[key] = val
}

func GetHash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// HandlerGET — обработчик запроса.
func (c RWMap) HandlerGET(w http.ResponseWriter, r *http.Request) {

	urlHash := chi.URLParam(r, "hash")
	log.Println(urlHash)

	if len(urlHash) != 32 {
		w.WriteHeader(500)
		return
	}

	var url string
	url = c.Get(urlHash)

	if len(url) > 0 {
		log.Println("result ", url)
		w.Header().Set("Location", url)
		w.WriteHeader(307)
		w.Write(nil)
	} else {
		w.WriteHeader(400)
	}

}

// HandlerPOST — обработчик запроса.
func (c RWMap) HandlerPOST(w http.ResponseWriter, r *http.Request) {

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
	c.Set(hash, address)

	log.Printf("Address = %s\n", address)
	w.WriteHeader(201)
	w.Write([]byte(Protocol + "://" + Host + ":" + Port + "/" + hash))

}
