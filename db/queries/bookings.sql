-- name: CreateBooking :one
INSERT INTO bookings (patient_id, patient_note, payment_id, total_cost, appointment_date)
VALUES ($1, $2, $3, $4, $5) RETURNING *;
