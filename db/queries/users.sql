-- name: CreateUser :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetPatient :one
SELECT *
FROM users
WHERE id = $1
  AND role = 'Patient';

-- name: UpdateUser :one
UPDATE users
SET full_name = $3,
    email = $4,
    phone_number = $5
WHERE id = $1 AND role = $2
RETURNING *;