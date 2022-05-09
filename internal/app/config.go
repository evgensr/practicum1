package app

import "github.com/evgensr/practicum1/internal/line"

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
	sessionName     = "practicum"
	ctxKeyUser  key = iota
)

type request struct {
	URL string `json:"url" valid:"url"`
}

type response struct {
	URL string `json:"result" valid:"url"`
}

type Line = line.Line

type Storage interface {
	Get(key string) string
	Set(url string, short string, user string) error
	Delete(key string) error
	GetByUser(key string) []Line
	Debug()
}
