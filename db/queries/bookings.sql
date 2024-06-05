-- name: CreateBooking :one
INSERT INTO bookings (patient_id, patient_note, payment_id)
VALUES ($1, $2, $3) RETURNING *;
