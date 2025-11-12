-- name: CreateTask :one
INSERT INTO tasks (title, description, project_id, assigned_to , status , updated_at)
VALUES ($1, $2, $3, $4, $5 ,$6)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks ORDER BY id LIMIT $1 OFFSET $2;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = $2
WHERE id = $1
RETURNING *;
