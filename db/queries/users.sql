-- name: CreateUser :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role, date_of_birth, gender)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
  AND deleted_at IS NULL;

-- name: GetUserByEmailForLogin :one
SELECT *
FROM users
WHERE email = $1
  AND deleted_at IS NULL;

-- name: GetPatient :one
SELECT *
FROM users
WHERE id = $1
  AND role = 'Patient';

-- name: UpdateUser :one
UPDATE users
SET full_name     = $2,
    email         = $3,
    phone_number  = $4,
    date_of_birth = $5,
    gender        = $6
WHERE id = $1 RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1;

-- name: IsEmailExists :one
SELECT EXISTS(SELECT 1
              FROM users
              WHERE email = $1
                AND deleted_at IS NULL) AS exists;

-- name: IsPhoneNumberExists :one
SELECT EXISTS(SELECT 1
              FROM users
              WHERE phone_number = $1
                AND deleted_at IS NULL) AS exists;
