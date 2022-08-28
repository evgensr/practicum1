package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/signal"
	"syscall"

	"github.com/evgensr/practicum1/internal/app"
	grpchandler "github.com/evgensr/practicum1/internal/grpc_handler"
	"github.com/evgensr/practicum1/internal/pb"
	"google.golang.org/grpc"
)

var buildVersion string = "N/A" // application version
var buildDate string = "N/A"    // application data
var buildCommit string = "N/A"  // commit id

type (
	Urls  int // количество сокращённых URL в сервисе
	Users int // количество пользователей в сервисе
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf := app.NewConfig()

	conf.Init()
	flag.Parse()

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	if len(conf.ConfigFile) > 0 {
		// Read and parse JSON file if flag -c with value exists
		jsonFileData, err := ioutil.ReadFile(conf.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		if err = json.Unmarshal(jsonFileData, &conf); err != nil {
			log.Fatal(err)
		}
	}

	server := app.New(&conf)

	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		server.GetLog().Fatal(err)
	}

	// создаём gRPC-сервер без зарегистрированной службы
	ss := grpc.NewServer()
	grpcHandler := grpchandler.NewGRPCHandler(server.GetStore())
	// регистрируем сервис

	pb.RegisterURLServer(ss, grpcHandler)

	go func() {

		server.GetLog().Info("Start gRPC server")
		// получаем запрос gRPC
		if err := ss.Serve(listen); err != nil {
			server.GetLog().Fatal(err)
		}
	}()

	if err := server.Start(ctx); err != nil {
		server.GetLog().Fatal(err)
	}

}
