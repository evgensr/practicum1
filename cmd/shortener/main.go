package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/app"
	"log"
	"os"
)

func main() {

	conf := app.NewConfig()
	err := env.Parse(&conf)

	if err != nil {
		log.Printf("[ERROR] failed to parse flags: %v", err)
		os.Exit(1)
	}

	flag.StringVar(&conf.ServerAddress, "a", "0.0.0.0:8080", "SERVER_ADDRESS")
	flag.StringVar(&conf.BaseURL, "b", "http://localhost:8080/", "BASE_URL")
	flag.StringVar(&conf.FileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	server := app.New(&conf)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
