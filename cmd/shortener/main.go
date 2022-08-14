package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/evgensr/practicum1/internal/app"
)

var buildVersion string = "N/A" // application version
var buildDate string = "N/A"    // application data
var buildCommit string = "N/A"  // commit id

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf := app.NewConfig()

	conf.Init()
	flag.Parse()

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	server := app.New(&conf)

	if err := server.Start(ctx); err != nil {
		log.Fatal(err)
	}

}
