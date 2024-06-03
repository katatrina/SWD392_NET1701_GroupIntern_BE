-- name: CreateAppointment :exec
INSERT INTO appointments (booking_id, schedule_id, customer_id)
VALUES ($1, $2, $3);