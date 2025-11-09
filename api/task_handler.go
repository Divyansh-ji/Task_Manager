package api

import (
	"net/http"
	db "projectmanager/db/gen"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ProjectID   int64  `json:"project_id" binding:"required"`
	//Status      string    `json:"status"`
	AssignedTo  any `json:"assigned_to" binding:"required"`
	//CreatedAt   time.Time `json:"createdat"`
}

func (server *Server) createTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var assignedToInt int
	switch v := req.AssignedTo.(type) {
	case float64:
		assignedToInt = int(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "assigned_to must be an integer"})
			return
		}
		assignedToInt = i
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "assigned_to must be an integer"})
		return
	}

	arg := db.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
		ProjectID:   int32(req.ProjectID),
		AssignedTo:  int32(assignedToInt),
	}
	task, err := server.store.CreateTask(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}
