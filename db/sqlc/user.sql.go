// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email,
  encrypted_password
) VALUES (
  $1, $2
)
RETURNING id, email, encrypted_password, created_at, updated_at
`

type CreateUserParams struct {
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.EncryptedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, encrypted_password, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET encrypted_password = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, email, encrypted_password, created_at, updated_at
`

type UpdateUserParams struct {
	ID                int32  `json:"id"`
	EncryptedPassword string `json:"encrypted_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.EncryptedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
