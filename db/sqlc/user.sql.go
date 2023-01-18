// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
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
    last_accessed_at,
    password_changed_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, username, first_name, last_name, hashed_password, email, user_type, active, created_at, last_accessed_at, password_changed_at
`

type CreateUserParams struct {
	ID                string    `json:"id"`
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	UserType          int32     `json:"user_type"`
	Active            int32     `json:"active"`
	CreatedAt         time.Time `json:"created_at"`
	LastAccessedAt    time.Time `json:"last_accessed_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.HashedPassword,
		arg.Email,
		arg.UserType,
		arg.Active,
		arg.CreatedAt,
		arg.LastAccessedAt,
		arg.PasswordChangedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.HashedPassword,
		&i.Email,
		&i.UserType,
		&i.Active,
		&i.CreatedAt,
		&i.LastAccessedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, first_name, last_name, hashed_password, email, user_type, active, created_at, last_accessed_at, password_changed_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.HashedPassword,
		&i.Email,
		&i.UserType,
		&i.Active,
		&i.CreatedAt,
		&i.LastAccessedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}
