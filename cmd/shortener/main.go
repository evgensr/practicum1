package main

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/app"
	"github.com/gorilla/sessions"

	"log"
)

var buildVersion string = "N/A" // application version
var buildDate string = "N/A"    // application data
var buildCommit string = "N/A"  // commit id

func main() {

	conf := app.NewConfig()
	err := env.Parse(&conf)

	// spew.Dump(conf)

	if err != nil {
		log.Fatalf("[ERROR] failed to parse flags: %v", err)
	}

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

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
