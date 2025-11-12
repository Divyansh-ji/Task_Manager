package api

import (
	"database/sql"
	"net/http"
	db "projectmanager/db/gen"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ProjectID   int64  `json:"project_id" binding:"required"`
	AssignedTo  int64  `json:"assigned_to"` // optional — can be null
	Status      string `json:"status"`      // optional — default will be set if empty
}

func (server *Server) createTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
		ProjectID:   int32(req.ProjectID),
		AssignedTo: sql.NullInt32{
			Int32: int32(req.AssignedTo),
			Valid: req.AssignedTo != 0, // valid only if non-zero user ID
		},
		Status: sql.NullString{
			String: func() string {
				if req.Status == "" {
					return "pending"
				}
				return req.Status
			}(),
			Valid: true,
		},
	}

	task, err := server.store.CreateTask(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
