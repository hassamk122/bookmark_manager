-- name: CreateUser :one
INSERT INTO users(username, email, password)
VALUES ($1, $2, $3)
RETURNING public_id, username, email, created_at, updated_at;

-- name: GetUserByID :one
SELECT public_id, username, email, created_at, updated_at
FROM users
WHERE public_id =$1;

-- name: GetUserByEmail :one
SELECT public_id, username, email, created_at, updated_at
FROM users
WHERE email = $1;


-- name: GetUserByEmailIncludePassword :one
SELECT public_id, username, email, password, created_at, updated_at
FROM users
WHERE email = $1;
