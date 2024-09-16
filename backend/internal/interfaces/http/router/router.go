package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamzidan/todolist/internal/interfaces/http/handler"
)

func Setup(h *handler.Task) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})

	r.GET("/tasks", h.GetListTask)
	r.GET("/tasks/:id", h.GetATask)
	r.POST("/tasks", h.CreateTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

	return r
}
