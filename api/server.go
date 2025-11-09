package api

import (
	"database/sql"

	db "projectmanager/db/gen"

	"github.com/gin-gonic/gin"
)

type Store interface {
	CreateTask(c *gin.Context, arg db.CreateTaskParams) (db.Task, error)
}

type queriesStore struct {
	q *db.Queries
}

func (s *queriesStore) CreateTask(c *gin.Context, arg db.CreateTaskParams) (db.Task, error) {
	return s.q.CreateTask(c.Request.Context(), arg)
}

type Server struct {
	store  Store
	router *gin.Engine
}

func NewServer(sqlDB *sql.DB) *Server {
	server := &Server{
		store:  &queriesStore{q: db.New(sqlDB)},
		router: gin.Default(),
	}
	server.router.POST("/tasks", server.createTask)
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
