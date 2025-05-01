-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE true
  AND email = @email
  AND deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (username, email, password, created_at, updated_at)
VALUES (@username, @email, @password, @created_at, @updated_at)
RETURNING *;
