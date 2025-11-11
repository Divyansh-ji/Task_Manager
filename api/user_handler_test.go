package api

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	db "projectmanager/db/gen"
	"testing"

	"github.com/gin-gonic/gin"
)

// --- Mock Store ---

func (m *mockStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	if arg.Email == "" {
		return db.User{}, errors.New("email is required")
	}
	return db.User{
		ID:        1,
		Email:     arg.Email,
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Password:  arg.Password,
	}, nil
}

func (m *mockStore) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	if id == 0 {
		return db.User{}, errors.New("invalid id")
	}

	return db.User{
		ID:        id,
		FirstName: "Divyansh",
		LastName:  "Tiwari",
		Email:     "DivyanshTiwary01@gmail.com",
		Password:  "hashedpassword",
	}, nil
}

// --- Actual Test ---
func TestUserHandler_Smoke(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := &Server{
		store:  &mockStore{},
		router: gin.Default(),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	body := `{
		"email": "DivyanshTiwary01@gmail.com",
		"first_name": "Divyansh",
		"last_name": "Tiwari",
		"password": "123456"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	server.CreateUsers(c)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}
}
