// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role, date_of_birth, gender)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
`

type CreateUserParams struct {
	FullName       string    `json:"full_name"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Role           string    `json:"role"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FullName,
		arg.HashedPassword,
		arg.Email,
		arg.PhoneNumber,
		arg.Role,
		arg.DateOfBirth,
		arg.Gender,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPatient = `-- name: GetPatient :one
SELECT id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
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
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmailForLogin = `-- name: GetUserByEmailForLogin :one
SELECT id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
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
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
FROM users
WHERE id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const isEmailExists = `-- name: IsEmailExists :one
SELECT EXISTS(SELECT 1
              FROM users
              WHERE email = $1
                AND deleted_at IS NULL) AS exists
`

func (q *Queries) IsEmailExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isEmailExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isPhoneNumberExists = `-- name: IsPhoneNumberExists :one
SELECT EXISTS(SELECT 1
              FROM users
              WHERE phone_number = $1
                AND deleted_at IS NULL) AS exists
`

func (q *Queries) IsPhoneNumberExists(ctx context.Context, phoneNumber string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isPhoneNumberExists, phoneNumber)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listPatientsByName = `-- name: ListPatientsByName :many
SELECT id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
FROM users
WHERE role = 'Patient'
  AND full_name ILIKE '%' || $1::text || '%'
  AND deleted_at IS NULL
ORDER BY created_at DESC
`

func (q *Queries) ListPatientsByName(ctx context.Context, fullName string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listPatientsByName, fullName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.HashedPassword,
			&i.Email,
			&i.PhoneNumber,
			&i.DateOfBirth,
			&i.Gender,
			&i.Role,
			&i.DeletedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name     = $2,
    email         = $3,
    phone_number  = $4,
    date_of_birth = $5,
    gender        = $6
WHERE id = $1 RETURNING id, full_name, hashed_password, email, phone_number, date_of_birth, gender, role, deleted_at, created_at
`

type UpdateUserParams struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.PhoneNumber,
		arg.DateOfBirth,
		arg.Gender,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	return err
}
