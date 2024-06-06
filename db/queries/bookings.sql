-- name: CreateExaminationBooking :one
INSERT INTO bookings (patient_id, patient_note, payment_id, total_cost, appointment_date, type)
VALUES ($1, $2, $3, $4, $5, 'examination') RETURNING *;

-- name: ListExaminationBookings :many
SELECT *
FROM bookings
WHERE patient_id = $1
  AND type = 'examination';
