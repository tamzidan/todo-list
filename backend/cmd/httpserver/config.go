package main

import (
	"time"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"github.com/tamzidan/todolist/internal/interfaces/http"
)

type Config struct {
	AppName     string    `env:"APP_NAME,required"`
	AppVersion  string    `env:"APP_VERSION,required"`
	AppLogLevel log.Level `env:"APP_LOG_LEVEL,required"`

	Host         string        `env:"HTTP_HOST,required"`
	Port         int           `env:"HTTP_PORT,required"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,required"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,required"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,required"`

	SQLite3DBPathName string `env:"SQLITE3_DB_PATHNAME,required"`
}

func (c *Config) GetHTTPConfig() http.Config {
	return http.Config{
		Host:         c.Host,
		Port:         c.Port,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		IdleTimeout:  c.IdleTimeout,
	}
}

func (c *Config) GetSQLite3Config() http.SQLite3Config {
	return http.SQLite3Config{
		DBPathName: c.SQLite3DBPathName,
	}
}

func mustLoadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("unable to parse configuration")
	}

	return cfg
}
