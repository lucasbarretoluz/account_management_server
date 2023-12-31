// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_name,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING id_user, user_name, hashed_password, full_name, email, create_at
`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.IDUser,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.CreateAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id_user, user_name, hashed_password, full_name, email, create_at FROM users
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.IDUser,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.CreateAt,
	)
	return i, err
}
