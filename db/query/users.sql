-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
