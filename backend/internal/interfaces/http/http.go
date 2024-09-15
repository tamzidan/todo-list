package http

import (
	"context"
	"fmt"
	nethttp "net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tamzidan/todolist/internal/interfaces/http/handler"
	"github.com/tamzidan/todolist/internal/interfaces/http/router"
	"github.com/tamzidan/todolist/internal/repository"
)

type SQLite3Config struct {
	DBPathName      string
	MigrationFolder string
}

type Config struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Server struct {
	cfg Config

	srv *nethttp.Server
}

func New(c Config, sqlite3Conf SQLite3Config) (*Server, error) {
	repo, err := repository.NewSQLite3Storage(sqlite3Conf.DBPathName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize sqlite connection")
	}

	h := handler.New(repo)
	r := router.Setup(h)

	srv := &nethttp.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Host, c.Port),
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		IdleTimeout:  c.IdleTimeout,
		Handler:      r,
	}

	return &Server{
		cfg: c,
		srv: srv,
	}, nil
}

// Run starts HTTP server as configured.
// This function is blocking, until the error is returned.
func (s *Server) Run() error {
	err := s.srv.ListenAndServe()
	return errors.Wrap(err, "running http server")
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.srv.Shutdown(ctx)
	return errors.Wrap(err, "stopping http server")
}
