package app

import (
	"github.com/evgensr/practicum1/internal/store/pg"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL         string `env:"BASE_URL,required" envDefault:"http://localhost:8080/"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	SessionKey      string `env:"SESSION_KEY" envDefault:"SESSION_KEY"`
	DatabaseDSN     string `env:"DATABASE_DSN,required" envDefault:"host=localhost user=postgres password=postgres dbname=restapi sslmode=disable"`
}

func NewConfig() Config {
	return Config{}
}

const (
	sessionName = "practicum"
	ctxKeyUser
)

type request struct {
	URL string `json:"url" valid:"url"`
}

type response struct {
	URL string `json:"response" valid:"url"`
}

type Storage interface {
	Get(key string) string
	Set(url string, short string, user string)
	Delete(key string) error
	GetByUser(key string) []pg.Line
	Debug()
}

//// Url структура Mutex
//type Url struct {
//	ShortUrl    string `json:"short_url"`
//	OriginalUrl string `json:"original_url"`
//}
//
//type User struct {
//	sync.RWMutex
//	id  string
//	url []Url
//}
