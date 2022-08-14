package app

import (
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		conf := NewConfig()
		err := env.Parse(&conf)
		assert.NoError(t, err)
		assert.Equal(t, "0.0.0.0:8080", conf.ServerAddress)
		assert.Equal(t, "http://localhost:8080/", conf.BaseURL)
		assert.Equal(t, "debug", conf.LogLevel)
		assert.Empty(t, conf.FileStoragePath)
		assert.Empty(t, conf.DatabaseDSN)
	})

	//conf = NewConfig()
	//t.Run("setenv config", func(t *testing.T) {
	//	conf := NewConfig()
	//	err := os.Setenv("SERVER_ADDRESS", "127.0.0.1:8181")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("SERVER_ADDRESS")
	//		assert.NoError(t, err)
	//	}()
	//
	//	err = os.Setenv("BASE_URL", "http://127.0.0.1:8181/")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("BASE_URL")
	//		assert.NoError(t, err)
	//	}()
	//
	//	err = os.Setenv("FILE_STORAGE_PATH", "store.txt")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("FILE_STORAGE_PATH")
	//		assert.NoError(t, err)
	//	}()
	//
	//	err = os.Setenv("LOG_LEVEL", "info")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("LOG_LEVEL")
	//		assert.NoError(t, err)
	//	}()
	//
	//	err = os.Setenv("SESSION_KEY", "key")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("SESSION_KEY")
	//		assert.NoError(t, err)
	//	}()
	//
	//	err = os.Setenv("DATABASE_DSN", "localhost")
	//	assert.NoError(t, err)
	//	defer func() {
	//		err = os.Unsetenv("DATABASE_DSN")
	//		assert.NoError(t, err)
	//	}()
	//
	//	// conf := NewConfig()
	//	err = env.Parse(&conf)
	//
	//	assert.NoError(t, err)
	//
	//	assert.Equal(t, "127.0.0.1:8181", conf.ServerAddress)
	//	assert.Equal(t, "http://127.0.0.1:8181/", conf.BaseURL)
	//	assert.Equal(t, "store.txt", conf.FileStoragePath)
	//	assert.Equal(t, "info", conf.LogLevel)
	//	assert.Equal(t, "key", conf.SessionKey)
	//	assert.Equal(t, "localhost", conf.DatabaseDSN)
	//
	//})
}
