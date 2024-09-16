package repository

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type SQLite3Storage struct {
	db *sql.DB
}

func NewSQLite3Storage(dbPathName string) (*SQLite3Storage, error) {
	db, err := sql.Open("sqlite3", dbPathName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open SQLite3 connection")
	}

	return &SQLite3Storage{db: db}, nil
}

func (s *SQLite3Storage) Close() error {
	return errors.Wrap(s.db.Close(), "unable to close database connection")
}

func (s *SQLite3Storage) GetTask(ctx context.Context, id int64) (Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, description, checked_at, created_at, updated_at, deleted_at FROM tasks WHERE id = ?", id).
		Scan(&t.Id, &t.Name, &t.Description, &t.CheckedAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, ErrNotFound
		}

		return Task{}, errors.Wrap(err, "unable to execute query")
	}

	return t, nil
}

func (s *SQLite3Storage) GetListTask(ctx context.Context, page, limit int) ([]Task, error) {
	offset := (page * limit) - limit
	rows, err := s.db.Query("SELECT id, name, description, checked_at, created_at, updated_at, deleted_at FROM tasks ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.CheckedAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			log.WithError(err).Error("unable to scan rows")
			continue
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "there is an error in query result rows")
	}

	return tasks, nil
}

func (s *SQLite3Storage) InsertTask(ctx context.Context, t Task) (Task, error) {
	stmt, err := s.db.Prepare("INSERT INTO tasks(name, description, checked_at, created_at, updated_at, deleted_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Task{}, errors.Wrap(err, "unable to prepare sql statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Name, t.Description, t.CheckedAt, t.CreatedAt, t.UpdatedAt, t.DeletedAt)
	if err != nil {
		return Task{}, errors.Wrap(err, "unble to execute insert query")
	}

	return t, nil
}

func (s *SQLite3Storage) UppdateTask(ctx context.Context, id int64, t Task) (Task, error) {
	_, err := s.db.Exec("UPDATE tasks SET name = ?, description = ?, checked_at = ?, created_at = ?, updated_at = ? WHERE id = ?", t.Name, t.Description, t.CheckedAt, t.CreatedAt, t.UpdatedAt, id)
	if err != nil {
		return Task{}, errors.Wrap(err, "unable to execte update query")
	}

	t.Id = id // just to make sure it has an id

	return t, nil
}

func (s *SQLite3Storage) DeleteTask(ctx context.Context, id int64) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return errors.Wrap(err, "unable to delete tasks")
	}
	return nil
}
