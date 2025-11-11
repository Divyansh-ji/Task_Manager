package api

import (
	"database/sql"
	"net/http"
	db "projectmanager/db/gen"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserParams struct {
	ID        int32
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt sql.NullTime
}

func (server *Server) CreateUsers(c *gin.Context) {
	var req UserParams
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	arg := db.CreateUserParams{
		ID:        req.ID,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}
	user, err := server.store.CreateUser(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user)

}
func (s *Server) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	id := int32(id64) // convert to int32

	user, err := s.store.GetUserByID(c, id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
