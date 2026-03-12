-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id)
VALUES ($1, $2)
RETURNING token_id, token, user_id, created_at, updated_at;

-- name: GetTokenByID :one
SELECT token_id, token, user_id, created_at, updated_at
FROM refresh_tokens
WHERE token_id = $1;

-- name: GetTokenByTokenString :one
SELECT token_id, token, user_id, created_at, updated_at
FROM refresh_tokens
WHERE token = $1;

-- name: DeleteTokenByUserID :exec
DELETE FROM refresh_tokens
WHERE user_id = $1;