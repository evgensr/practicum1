package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/app"
	"github.com/gorilla/sessions"
	"log"
	"os"
)

func main() {

	conf := app.NewConfig()
	err := env.Parse(&conf)

	// spew.Dump(conf)

	if err != nil {
		log.Printf("[ERROR] failed to parse flags: %v", err)
		os.Exit(1)
	}

	flag.StringVar(&conf.ServerAddress, "a", os.Getenv("SERVER_ADDRESS"), "SERVER_ADDRESS")
	flag.StringVar(&conf.BaseURL, "b", os.Getenv("BASE_URL"), "BASE_URL")
	flag.StringVar(&conf.FileStoragePath, "f", os.Getenv("FILE_STORAGE_PATH"), "FILE_STORAGE_PATH")
	flag.StringVar(&conf.DatabaseDSN, "d", os.Getenv("DATABASE_DSN"), "DATABASE_DSN")

	flag.Parse()

	sessionStore := sessions.NewCookieStore([]byte(conf.SessionKey))

	server := app.New(&conf, sessionStore)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
