package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := &Server{
		store: &mockStore{},
	}

	// Create a fake HTTP request
	body := `{
		"title": "Test Task",
		"description": "Testing mock",
		"project_id": 1,
		"assigned_to": 2
	}`

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	// Call the handler
	server.createTask(c)

	// Validate response
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}
}
