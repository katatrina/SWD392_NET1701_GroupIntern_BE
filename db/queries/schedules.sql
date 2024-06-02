-- name: UpdateExaminationSchedule :exec
UPDATE examination_schedules
SET booking_id  = coalesce(sqlc.narg('booking_id'), name),
    customer_id = coalesce(sqlc.narg('customer_id'), bio)
WHERE id = sqlc.arg('id');
