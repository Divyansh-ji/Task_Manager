package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	db "projectmanager/db/gen"
	"strconv"
	"testing" //this is the package already provided by the go for testing

	"github.com/gin-gonic/gin"
)

type mockStore struct{}

func (m *mockStore) CreateTask(c *gin.Context, arg db.CreateTaskParams) (db.Task, error) {
	if arg.Title == "" {
		return db.Task{}, errors.New("title is requried")
	}
	return db.Task{
		ID:          1,
		Title:       arg.Title,
		Description: arg.Description,
		ProjectID:   arg.ProjectID,
		AssignedTo:  strconv.Itoa(int(arg.AssignedTo)),
	}, nil
}

func (m *mockStore) GetTask(task *CreateTaskRequest, id int32) *mockError {
	if id == 0 {
		return &mockError{err: errors.New("invalid id")}
	}
	task.Title = "Mock Task"
	task.ProjectID = 1
	task.AssignedTo = "2"
	return &mockError{}
}

type mockError struct {
	err error
}

func (e *mockError) Error() error {
	return e.err
}
func TestCreateTask(t *testing.T) {
	server := &Server{store: &mockStore,
		router: gin.Default()} // use fake DB instead of real one

	recorder := httptest.NewRecorder()      // fake HTTP response recorder
	c, _ := gin.CreateTestContext(recorder) // create a fake Gin context

	// send a fake JSON request
	body := `{
		"title": "Test Task",
		"description": "Testing mock",
		"project_id": 1,
		"assigned_to": 2
	}`
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	server.createTask(c) // call your real handler with the fake request

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}
}
