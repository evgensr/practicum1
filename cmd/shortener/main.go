package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/app"
	"github.com/gorilla/sessions"
	"log"
	"os"
)

var version = "0.0.1"

func main() {

	conf := app.NewConfig()
	err := env.Parse(&conf)

	// spew.Dump(conf)

	if err != nil {
		log.Printf("[ERROR] failed to parse flags: %v", err)
		os.Exit(1)
	}

	flag.StringVar(&conf.ServerAddress, "a", conf.ServerAddress, "SERVER_ADDRESS")
	flag.StringVar(&conf.BaseURL, "b", conf.BaseURL, "BASE_URL")
	flag.StringVar(&conf.FileStoragePath, "f", conf.FileStoragePath, "FILE_STORAGE_PATH")
	flag.StringVar(&conf.DatabaseDSN, "d", conf.DatabaseDSN, "DATABASE_DSN")

	flag.Parse()

	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))

	server := app.New(&conf, sessionStore)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
