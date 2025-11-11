package api

import (
	"context"
	"database/sql"

	db "projectmanager/db/gen"

	"github.com/gin-gonic/gin"
)

// Interface for all DB operations
type Store interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error)
}

// Struct that implements Store (wraps db.Queries)
type queriesStore struct {
	q *db.Queries
}

func (s *queriesStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return s.q.CreateUser(ctx, arg)
}

func (s *queriesStore) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return s.q.GetUserByID(ctx, id)
}

func (s *queriesStore) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	return s.q.CreateTask(ctx, arg)
}

// --- Server struct ---
type Server struct {
	store  Store
	router *gin.Engine
}

// --- Initialize Server ---
func NewServer(sqlDB *sql.DB) *Server {
	server := &Server{
		store:  &queriesStore{q: db.New(sqlDB)},
		router: gin.Default(),
	}

	// ✅ just route to your existing handlers
	server.router.POST("/users", server.CreateUsers)
	server.router.POST("/tasks", server.createTask)
	server.router.GET("/users", server.GetUserByID)

	return server
}

// --- Start server ---
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
