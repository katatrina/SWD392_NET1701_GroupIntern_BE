// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, 'customer') RETURNING id, full_name, hashed_password, email, phone_number, role, created_at
`

type CreateCustomerParams struct {
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.FullName,
		arg.HashedPassword,
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
		&i.CreatedAt,
	)
	return i, err
}

const getCustomerByEmail = `-- name: GetCustomerByEmail :one
SELECT id, full_name, hashed_password, email, phone_number, role, created_at FROM users WHERE email = $1 AND role = 'customer'
`

func (q *Queries) GetCustomerByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}