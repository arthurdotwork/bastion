-- name: CreateAccessToken :one
INSERT INTO access_tokens (id, user_id, token_identifier, issued_at, expires_at)
VALUES (@id, @user_id, @token_identifier, @issued_at, @expires_at)
RETURNING *;
