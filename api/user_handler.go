package api

import (
	"database/sql"
	"net/http"
	utils "projectmanager/auth"
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
func (s *Server) RegisterUser(c *gin.Context) {
	var req db.CreateUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashed, _ := utils.HashPassword(req.Password)
	req.Password = hashed
	user, err := s.store.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}
func (s *Server) Login(c *gin.Context) {
	var req struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Step 1: Bind request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Fetch user from DB by email
	user, err := s.store.GetUserByID(c.Request.Context(), int32(req.ID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Step 3: Compare passwords (bcrypt)
	if !utils.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Step 4: Generate JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// Step 5: Respond with token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
