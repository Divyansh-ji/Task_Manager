package api

import (
	"net/http"
	db "projectmanager/db/gen"

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
	arg := db.CreateProjectParams{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     int32(req.OwnerID),
	}
	project, err := s.store.CreateProject(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)

}
