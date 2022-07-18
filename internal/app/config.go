package app

import (
	"github.com/evgensr/practicum1/internal/store"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL         string `env:"BASE_URL,required" envDefault:"http://localhost:8080/"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	SessionKey      string `env:"SESSION_KEY" envDefault:"SESSION_KEY"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

func NewConfig() Config {
	return Config{}
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

type Line = store.Line

type Storage interface {
	Get(key string) (Line, error)
	Set(Line) error
	Delete([]Line) error
	GetByUser(key string) []Line
}
