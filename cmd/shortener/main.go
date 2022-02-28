package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/cmd/shortener/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {

	// Init
	counter := handler.NewRWMap()

	var cfg handler.Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg.ServerAddress)
	log.Println("start server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{hash}", counter.HandlerGET)
	r.Post("/", counter.HandlerPOST)
	r.Post("/api/shorten", counter.HandlerPOST2)
	err = http.ListenAndServe(cfg.ServerAddress, r)

	if err != nil {
		log.Fatal(err)
	}

}
