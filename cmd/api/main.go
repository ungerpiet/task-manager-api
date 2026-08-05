package main

import (
	"context"
	"go-task-api/internal/database"
	"go-task-api/internal/handler"
	"go-task-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	conn := database.ConnectDB()
	defer conn.Close(context.Background())

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Preflight-Anfragen direkt beantworten
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.Static("/web", "./web")
	r.StaticFile("/", "./web/index.html")

	repo := repository.NewTaskRepository(conn)
	handler := handler.NewTaskHandler(repo)

	r.GET("/tasks", handler.GetTasks)
	r.POST("/tasks", handler.CreateTask)

	r.GET("/tasks/:id", handler.GetTaskByID)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	r.Run(":8080") // Startet Server auf localhost:8080
}
