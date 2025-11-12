package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     int64  `json:"owner_id" binding:"required"`
}

func (s *Server) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid req",
			})
			return
		}

	}

}
