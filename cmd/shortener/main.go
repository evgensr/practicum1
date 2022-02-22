package main

import (
	"github.com/evgensr/practicum1/cmd/shortener/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {

	// Init
	counter := handler.NewRWMap()

	log.Println("start server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{hash}", counter.HandlerGET)
	r.Post("/", counter.HandlerPOST)
	err := http.ListenAndServe(":"+handler.Port, r)

	if err != nil {
		log.Fatal(err)
	}

}
