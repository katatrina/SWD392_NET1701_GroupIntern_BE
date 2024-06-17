-- name: CreatePayment :one
INSERT INTO payments (name)
VALUES ($1) RETURNING *;

-- name: ListPayments :many
SELECT * FROM payments;
