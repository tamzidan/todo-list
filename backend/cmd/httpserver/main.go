package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tamzidan/todolist/internal/interfaces/http"
	"github.com/tamzidan/todolist/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger.Setup(log.DebugLevel)

	log.Println("loading configurations...")

	cfg := mustLoadConfig()
	httpConfig := http.Config{
		Host:         cfg.Host,
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	sqlite3Config := http.SQLite3Config{
		DBPathName:      cfg.SQLite3DBPathName,
		MigrationFolder: cfg.DBMigrationFolder,
	}

	log.Println("migration database...")

	mustMigrate(sqlite3Config.DBPathName, sqlite3Config.MigrationFolder)

	server, err := http.New(httpConfig, sqlite3Config)
	if err != nil {
		log.WithError(err).Fatal("unable to initialize http server")
	}

	log.WithField("config", httpConfig).Info("starting http server")
	go func(s *http.Server) {
		if err := s.Run(); err != nil {
			log.WithError(err).Error("unable to run http server")
			return
		}
	}(server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutdown signal received. shutting down server.")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.WithError(err).Error("unable to shutdown http server")
		return
	}
}

func mustMigrate(dbFile, migrationFolder string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("unable to open sqlite3 connection: %s", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("unable to initialize")
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationFolder), dbFile, driver)
	if err != nil {
		log.Fatalf("unable to initialize database instance: %s", err)
	}

	if err = m.Up(); err != nil {
		if errDown := m.Down(); errDown != nil {
			log.Fatalf("unable to migrate down: %s", errDown)
		}

		log.Fatalf("unable to migrate up: %s", err)
	}
}
