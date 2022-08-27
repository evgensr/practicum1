package app

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/evgensr/practicum1/internal/store"
)

type Config struct {
	ServerAddress   string `json:"server_address" env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL         string `json:"base_url" env:"BASE_URL,required" envDefault:"http://localhost:8080/"`
	FileStoragePath string `json:"file_storage_path" env:"FILE_STORAGE_PATH"`
	LogLevel        string `json:"log_level" env:"LOG_LEVEL" envDefault:"debug"`
	SessionKey      string `env:"SESSION_KEY" envDefault:"SESSION_KEY"`
	DatabaseDSN     string `json:"database_dsn" env:"DATABASE_DSN"`
	EnableHTTPS     bool   `json:"enable_https" env:"ENABLE_HTTPS" envDefault:"false"`
	ConfigFile      string `env:"CONFIG"`
	TrustedSubnet   string `json:"trusted_subnet" env:"TRUSTED_SUBNET"`
}

func NewConfig() Config {

	var conf Config

	err := env.Parse(&conf)

	if err != nil {
		log.Fatalf("[ERROR] failed to parse flags: %v", err)
	}

	//flag.StringVar(&conf.ServerAddress, "a", conf.ServerAddress, "SERVER_ADDRESS")
	//flag.StringVar(&conf.BaseURL, "b", conf.BaseURL, "BASE_URL")
	//flag.StringVar(&conf.FileStoragePath, "f", conf.FileStoragePath, "FILE_STORAGE_PATH")
	//flag.StringVar(&conf.DatabaseDSN, "d", conf.DatabaseDSN, "DATABASE_DSN")
	//flag.BoolVar(&conf.EnableHTTPS, "s", conf.EnableHTTPS, "ENABLE_HTTPS, files needed server.crt and server.key in the current directory")
	//flag.StringVar(&conf.ConfigFile, "c", conf.ConfigFile, "JSON config file")
	//
	//flag.Parse()

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

	return conf
}

type key int

const (
	sessionName     = "practicum" // cookies for the user
	ctxKeyUser  key = iota        // context id req
)

type request struct {
	URL string `json:"url" valid:"url"`
}

type response struct {
	URL string `json:"result" valid:"url"`
}

type (
	Line = store.Line
	// Urls  int // количество сокращённых URL в сервисе
	// Users int // количество пользователей в сервисе
)

type Storage interface {
	Get(key string) (Line, error)
	Set(Line) error
	Delete([]Line) error
	GetByUser(key string) []Line
	GetStats() (store.Urls, store.Users, error)
	Shutdown(ctx context.Context) error
}

func (c *Config) Init() {

	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "SERVER_ADDRESS")
	flag.StringVar(&c.BaseURL, "b", c.BaseURL, "BASE_URL")
	flag.StringVar(&c.FileStoragePath, "f", c.FileStoragePath, "FILE_STORAGE_PATH")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "DATABASE_DSN")
	flag.BoolVar(&c.EnableHTTPS, "s", c.EnableHTTPS, "ENABLE_HTTPS, files needed server.crt and server.key in the current directory")
	flag.StringVar(&c.ConfigFile, "c", c.ConfigFile, "JSON config file")

}
