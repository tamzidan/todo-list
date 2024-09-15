package main

import (
	"time"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
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

	DBMigrationFolder string `env:"DB_MIGRATION_FOLDER,required"`

	SQLite3DBPathName string `env:"SQLITE3_DB_PATHNAME,required"`
}

func mustLoadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("unable to parse configuration")
	}

	return cfg
}
