package api

import (
	"context"
	"database/sql"
	"errors"
	db "projectmanager/db/gen"
)

// --- Mock Store implementing only what we need ---
type mockStore struct{}

func (m *mockStore) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	if arg.Title == "" {
		return db.Task{}, errors.New("title is required")
	}
	return db.Task{
		ID:          1,
		Title:       arg.Title,
		Description: arg.Description,
		ProjectID:   arg.ProjectID,
		AssignedTo:  arg.AssignedTo,
	}, nil
}

// --- Optional: mock for GetTask (if needed elsewhere) ---
func (m *mockStore) GetTask(ctx context.Context, id int32) (db.Task, error) {
	if id == 0 {
		return db.Task{}, errors.New("invalid id")
	}
	return db.Task{
		ID:          id,
		Title:       "Sample Task",
		Description: "Mocked description",
		ProjectID:   1,
		AssignedTo:  sql.NullInt32{Int32: 2, Valid: true},
	}, nil
}
func (m *mockStore) CreateProject(ctx context.Context, arg db.CreateProjectParams) (db.Project, error) {
	return db.Project{
		ID:          1,
		Name:        arg.Name,
		Description: arg.Description,
		OwnerID:     arg.OwnerID,
	}, nil
}
