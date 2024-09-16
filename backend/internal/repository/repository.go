package repository

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type Task struct {
	Id          int64
	Name        string
	Description string
	CheckedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type TodoListStorage interface {
	GetListTask(ctx context.Context, page, limit int) ([]Task, error)
	GetTask(ctx context.Context, id int64) (Task, error)
	InsertTask(ctx context.Context, t Task) (Task, error)
	UppdateTask(ctx context.Context, id int64, t Task) (Task, error)
	DeleteTask(ctx context.Context, id int64) error
}
