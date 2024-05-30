-- name: CreateExaminationBooking :one
INSERT INTO bookings (type, customer_id, customer_reason, payment_id)
VALUES ('examination', $1, $2, $3) RETURNING *;
