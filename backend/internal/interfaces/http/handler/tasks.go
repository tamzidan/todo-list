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
	Error ErrorResponse      `json:"error"`
	Data  any                `json:"data"`
	Page  PaginationResponse `json:"page"`
}

type PaginationResponse struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type ErrorResponse struct {
	Code    int
	Message string
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
