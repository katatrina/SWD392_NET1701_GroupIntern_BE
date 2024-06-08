-- name: CreatePatient :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, 'patient') RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetPatient :one
SELECT *
FROM users
WHERE id = $1
  AND role = 'patient';
