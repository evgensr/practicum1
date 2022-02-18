package main

import (
	"github.com/evgensr/practicum1/cmd/shortener/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

// handlerURL удалил, он был написан на первом этапе, без использования фрейворка.

func main() {

	log.Println("start server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{hash}", handler.HandlerGET)
	r.Post("/", handler.HandlerPOST)
	err := http.ListenAndServe(":"+handler.Port, r)

	if err != nil {
		log.Fatal(err)
	}

}
