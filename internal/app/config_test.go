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

func TestConfig_Init(t *testing.T) {
	type fields struct {
		ServerAddress   string
		BaseURL         string
		FileStoragePath string
		LogLevel        string
		SessionKey      string
		DatabaseDSN     string
		EnableHTTPS     bool
		ConfigFile      string
		TrustedSubnet   string
		GrpcPort        string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "valid",
			fields: struct {
				ServerAddress   string
				BaseURL         string
				FileStoragePath string
				LogLevel        string
				SessionKey      string
				DatabaseDSN     string
				EnableHTTPS     bool
				ConfigFile      string
				TrustedSubnet   string
				GrpcPort        string
			}{ServerAddress: "0.0.0.0:8080",
				BaseURL:         "http://localhost:8080/",
				FileStoragePath: "file.db",
				LogLevel:        "debug",
				SessionKey:      "SESSION_KEY",
				DatabaseDSN:     "DatabaseDSN",
				EnableHTTPS:     false,
				ConfigFile:      "",
				TrustedSubnet:   "",
				GrpcPort:        ":3310"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				ServerAddress:   tt.fields.ServerAddress,
				BaseURL:         tt.fields.BaseURL,
				FileStoragePath: tt.fields.FileStoragePath,
				LogLevel:        tt.fields.LogLevel,
				SessionKey:      tt.fields.SessionKey,
				DatabaseDSN:     tt.fields.DatabaseDSN,
				EnableHTTPS:     tt.fields.EnableHTTPS,
				ConfigFile:      tt.fields.ConfigFile,
				TrustedSubnet:   tt.fields.TrustedSubnet,
				GrpcPort:        tt.fields.GrpcPort,
			}
			c.Init()
		})
	}
}

func TestNewConfig(t *testing.T) {
	conf := NewConfig()

	tests := []struct {
		name string
		want Config
	}{
		{
			name: "valid",
			want: conf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewConfig(), "NewConfig()")
		})
	}
}
