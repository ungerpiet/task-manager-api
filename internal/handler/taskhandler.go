package handler

import (
	"errors"
	"fmt"
	"go-task-api/internal/model"
	"go-task-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (r *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := r.repo.GetAll()
	if err != nil {
		fmt.Println("GetAll DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (r *TaskHandler) CreateTask(c *gin.Context) {
	var req model.Task

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("BindJSON Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := r.repo.Create(&req); err != nil {
		fmt.Println("Create DB Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (r *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	task, err := r.repo.GetByID(id)

	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("GetByID DB Error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not Found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (r *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req model.Task

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("BindJSON Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = r.repo.Update(&req, id)
	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("Update DB Error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	req.ID = id
	c.JSON(http.StatusOK, req)
}

func (r *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = r.repo.Delete(id)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
