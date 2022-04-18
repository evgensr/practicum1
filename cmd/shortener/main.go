package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/app"
	"log"
	"os"
)

func main() {

	// var cfg app.Config
	conf := app.NewConfig()
	// var cfg app.Config
	// spew.Dump(conf)
	// spew.Dump(cfg)
	err := env.Parse(&conf)
	// spew.Dump(conf)
	// log.Println(conf)
	if err != nil {
		log.Printf("[ERROR] failed to parse flags: %v", err)
		os.Exit(1)
	}

	// spew.Dump(cfg)

	server := app.New(&conf)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	//var cfg config.Config
	//err := env.Parse(&cfg)
	//if err != nil {
	//	log.Fatalf("[ERROR] can't make server, %v", err)
	//}
	//api := handlers.API{}
	//
	//log.Println(cfg.ServerAddress)
	//
	//r := chi.NewRouter()
	//r.Get("/{hash}", api.Get)
	//r.Post("/api/shorten", handlers.Post)
	//
	//err = http.ListenAndServe(cfg.ServerAddress, r)
	//
	//if err != nil {
	//	log.Fatalf("[ERROR] can't make server, %v", err)
	//}

}
