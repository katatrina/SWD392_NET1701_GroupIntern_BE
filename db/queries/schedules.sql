-- name: ListExaminationSchedulesByDateAndServiceCategory :many
SELECT s.id as schedule_id, s.type, s.start_time, s.end_time, u.full_name as dentist_name, r.name as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         JOIN examination_schedule_detail esd ON s.id = esd.schedule_id
WHERE s.start_time::date = sqlc.arg(date)::date
AND esd.service_category_id = sqlc.arg(service_category_id)
ORDER BY s.start_time ASC;

-- name: GetExaminationScheduleDetail :one
SELECT s.id        as schedule_id,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       sc.name     as service_category_name,
       sc.cost     as service_category_cost
FROM schedules s
         JOIN examination_schedule_detail sd ON s.id = sd.schedule_id
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         JOIN service_categories sc ON sd.service_category_id = sc.id
WHERE s.id = sqlc.arg(schedule_id);

-- name: CreateSchedule :one
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: CreateExaminationScheduleDetail :one
INSERT INTO examination_schedule_detail (schedule_id, service_category_id)
VALUES ($1, $2) RETURNING *;

-- name: UpdateExaminationScheduleSlotsRemaining :exec
UPDATE examination_schedule_detail
SET slots_remaining = slots_remaining - 1
WHERE schedule_id = $1;
