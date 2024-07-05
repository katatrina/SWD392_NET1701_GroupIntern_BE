-- name: CreateUser :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role, date_of_birth, gender)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

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