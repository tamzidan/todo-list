package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tamzidan/todolist/internal/interfaces/http"
	"github.com/tamzidan/todolist/pkg/logger"
)

func main() {
	fmt.Println("Hello world")
	log.Info("test...")

	logger.Setup(log.DebugLevel)

	// TODO Add migration script

	cfg := mustLoadConfig()
	httpConfig := cfg.GetHTTPConfig()
	sqlite3Config := cfg.GetSQLite3Config()
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
