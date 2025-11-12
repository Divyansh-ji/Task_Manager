-- name: CreateProject :one
INSERT INTO projects (name, description, owner_id)
VALUES (
  sqlc.arg(name),
  sqlc.arg(description),
  sqlc.arg(owner_id)
)
RETURNING *;
