-- name: CreateCustomer :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, 'customer') RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
