package handler

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tamzidan/todolist/internal/repository"
)

type Task struct {
	repo repository.TodoListStorage
}

func New(repo repository.TodoListStorage) *Task {
	return &Task{
		repo: repo,
	}
}

type Response struct {
	Error ErrorResponse      `json:"error,omitempty"`
	Data  any                `json:"data"`
	Page  PaginationResponse `json:"page"`
}

type PaginationResponse struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type CreateTaskReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (t *Task) CreateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var reqBody CreateTaskReq
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusBadRequest),
			},
		})
		return
	}

	_, err := t.repo.InsertTask(ctx, repository.Task{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		log.WithError(err).Error("unable to insert new task")
		c.JSON(http.StatusInternalServerError, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Data: map[string]string{},
		Page: PaginationResponse{},
	})
	return
}

type UpdateTaskReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (t *Task) UpdateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)
	if id < 1 {
		c.JSON(http.StatusBadRequest, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusBadRequest),
			},
		})
		return
	}

	var reqBody UpdateTaskReq
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusBadRequest),
			},
		})
		return
	}

	_, err := t.repo.UppdateTask(ctx, int64(id), repository.Task{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		log.WithError(err).Error("unable to update data in database")
		c.JSON(http.StatusInternalServerError, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: map[string]string{},
		Page: PaginationResponse{},
	})
	return
}

func (t *Task) DeleteTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)
	if id < 1 {
		c.JSON(http.StatusBadRequest, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusBadRequest),
			},
		})
		return
	}

	if err := t.repo.DeleteTask(ctx, int64(id)); err != nil {
		log.WithError(err).Error("unable to delete task in database")
		c.JSON(http.StatusInternalServerError, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: map[string]string{},
		Page: PaginationResponse{},
	})
	return
}

func (t *Task) GetATask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)
	if id < 1 {
		c.JSON(http.StatusBadRequest, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusBadRequest),
			},
		})
		return
	}

	task, err := t.repo.GetTask(ctx, int64(id))
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(http.StatusNotFound, Response{
			Error: ErrorResponse{
				Code:    1, // TODO We should have list of error code
				Message: "record not found",
			},
		})
		return
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		log.WithError(err).Error("unable to get a tasks from database")
		c.JSON(http.StatusInternalServerError, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: task,
		Page: PaginationResponse{},
	})
	return
}

func (t *Task) GetListTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	pageReq, ok := c.Get("page")
	if !ok {
		pageReq = 1
	}

	currPage := pageReq.(int)
	limit := 100 // by default, we don't accept value from req for now

	taskList, err := t.repo.GetListTask(ctx, currPage, limit)
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(http.StatusNotFound, Response{
			Error: ErrorResponse{
				Code:    1, // TODO We should have list of error code
				Message: "record not found",
			},
		})
		return
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		log.WithError(err).Error("unable to get list of tasks from database")
		c.JSON(http.StatusInternalServerError, Response{
			Error: ErrorResponse{
				Code:    1,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: taskList,
		Page: PaginationResponse{
			Next: nextURL(*c.Request.URL, currPage, limit, len(taskList)),
			Prev: prevURL(*c.Request.URL, currPage),
		},
	})
	return
}

func nextURL(currentURL url.URL, currPage, limit, dataLen int) string {
	if dataLen < limit {
		return ""
	}

	q := currentURL.Query()
	q.Set("page", strconv.Itoa(currPage+1))
	currentURL.RawQuery = q.Encode()

	return currentURL.String()
}

func prevURL(currentURL url.URL, currPage int) string {
	if currPage <= 1 {
		return ""
	}

	q := currentURL.Query()
	q.Set("page", strconv.Itoa(currPage-1))
	currentURL.RawQuery = q.Encode()

	return currentURL.String()
}
