// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, $5) RETURNING id, full_name, hashed_password, email, phone_number, role, deleted_at, created_at
`

type CreateUserParams struct {
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Role           string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FullName,
		arg.HashedPassword,
		arg.Email,
		arg.PhoneNumber,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPatient = `-- name: GetPatient :one
SELECT id, full_name, hashed_password, email, phone_number, role, deleted_at, created_at
FROM users
WHERE id = $1
  AND role = 'Patient'
`

func (q *Queries) GetPatient(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getPatient, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmailForLogin = `-- name: GetUserByEmailForLogin :one
SELECT id, full_name, hashed_password, email, phone_number, role, deleted_at, created_at
FROM users
WHERE email = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetUserByEmailForLogin(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailForLogin, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name = $3,
    email = $4,
    phone_number = $5
WHERE id = $1 AND role = $2
RETURNING id, full_name, hashed_password, email, phone_number, role, deleted_at, created_at
`

type UpdateUserParams struct {
	ID          int64  `json:"id"`
	Role        string `json:"role"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Role,
		arg.FullName,
		arg.Email,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}
