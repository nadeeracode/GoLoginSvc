-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    first_name,
    last_name,
    hashed_password,
    email,
    user_type,
    active,
    created_at,
    last_accessed_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;