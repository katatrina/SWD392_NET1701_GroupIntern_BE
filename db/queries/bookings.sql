-- name: CreateBooking :one
INSERT INTO bookings (patient_id, type, payment_status, payment_id, total_cost, appointment_date)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: ListExaminationBookings :many
SELECT *
FROM bookings
WHERE patient_id = $1
  AND type = 'Examination';
