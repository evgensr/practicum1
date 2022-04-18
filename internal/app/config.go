package app

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL         string `env:"BASE_URL,required" envDefault:"http://localhost:8080/"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() Config {
	return Config{}
}

type request struct {
	URL string `json:"url" valid:"url"`
}

type response struct {
	URL string `json:"response" valid:"url"`
}

type Storage interface {
	Get(key string) string
	Set(key string, val string)
	Delete(key string) error
	Debug()
}
