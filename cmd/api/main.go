package main

import (
	"context"
	"go-task-api/internal/database"
	"go-task-api/internal/handler"
	"go-task-api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	conn := database.ConnectDB()
	defer conn.Close(context.Background())

	repo := repository.NewTaskRepository(conn)
	handler := handler.NewTaskHandler(repo)

	r := gin.Default()

	r.GET("/tasks", handler.GetTasks)
	r.POST("/tasks", handler.CreateTask)

	r.GET("/tasks/:id", handler.GetTaskByID)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8080") // Startet Server auf localhost:8080
}
