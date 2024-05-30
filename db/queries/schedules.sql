-- name: UpdateExaminationSchedule :exec
UPDATE examination_schedules SET booking_id = $2 AND customer_id = $3 WHERE id = $1;
