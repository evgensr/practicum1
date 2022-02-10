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

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			log.Printf("ParseForm() err: %v", err)
			return
		}

		// address := r.FormValue("address")
		address := string(b)
		hash := getHash(address)
		log.Println(hash)
		mapURL[hash] = address
		log.Printf("Address = %s\n", address)
		w.WriteHeader(201)
		w.Write([]byte("http://localhost:" + port + "/" + hash))

		// mapURL[]
	case "GET":
		// hash := getHash(r.RequestURI)
		// url := getHash(r.RequestURI)
		// url := r.RequestURI
		// inputFmt := r.RequestURI[1:]
		//log.Println(inputFmt)
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
		// md5.Sum()
		// data := []byte(r.RequestURI)
		// h := md5.New()
		// h.Sum(data)
		// fmt.Println(time.RFC822Z, md5.Sum(data), h.Sum(data), h.Sum([]byte("/12")))
		// fmt.Println("GET")
	}

	// fmt.Println(r.Header)
	// fmt.Println(time.RFC822Z, r.RequestURI)

	// w.Write([]byte("<h1>Hello, World</h1>"))
}

func form(w http.ResponseWriter, r *http.Request) {

	log.Println("form")
	http.ServeFile(w, r, "form.html")
}

func getHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

	log.Println("start server")

	// маршрутизация запросов обработчику
	http.HandleFunc("/", handlerURL)
	http.HandleFunc("/create", form)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	// запуск сервера с адресом localhost, порт 8081
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
