package api

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	db "projectmanager/db/gen"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockStore struct{}

// --- Mock CreateTask ---
func (m *mockStore) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	if arg.Title == "" {
		return db.Task{}, errors.New("title is required")
	}
	return db.Task{
		ID:          1,
		Title:       arg.Title,
		Description: arg.Description,
		ProjectID:   arg.ProjectID,
		AssignedTo:  strconv.Itoa(int(arg.AssignedTo)),
	}, nil
}

// --- Mock GetTask ---
func (m *mockStore) GetTask(ctx context.Context, id int32) (db.Task, error) {
	if id == 0 {
		return db.Task{}, errors.New("invalid id")
	}

	// return a fake task for testing
	return db.Task{
		ID:          id,
		Title:       "Sample Task",
		Description: "Mocked description",
		ProjectID:   1,
		AssignedTo:  "2",
	}, nil
}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := &Server{
		store:  &mockStore{}, // fake DB
		router: gin.Default(),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	// fake request body
	body := `{
		"title": "Test Task",
		"description": "Testing mock",
		"project_id": 1,
		"assigned_to": 2
	}`
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	server.createTask(c) // call the actual handler

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}
}
