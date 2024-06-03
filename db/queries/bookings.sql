-- name: CreateBooking :one
INSERT INTO bookings (customer_id, customer_note, payment_id)
VALUES ($1, $2, $3) RETURNING *;
