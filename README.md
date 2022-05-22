# go-musthave-shortener-tpl
Шаблон репозитория для практического трека «Go в веб-разработке».

# Начало работы

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.
2. В корне репозитория выполните команду `go mod init <name>` (где `<name>` - адрес вашего репозитория на GitHub без префикса `https://`) для создания модуля.

# Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона выполните следующую команду:

```
git remote add -m main template https://github.com/yandex-praktikum/go-musthave-shortener-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/main .github
```

Затем добавьте полученные изменения в свой репозиторий.


### Конфиг
```
type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL         string `env:"BASE_URL,required" envDefault:"http://localhost:8080/"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	SessionKey      string `env:"SESSION_KEY" envDefault:"SESSION_KEY"`
	DatabaseDSN     string `env:"DATABASE_DSN" envDefault:"host=localhost user=postgres password=postgres dbname=restapi sslmode=disable"`
}
```