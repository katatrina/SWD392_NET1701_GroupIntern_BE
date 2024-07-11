-- name: ListAvailableExaminationSchedulesByDateForPatient :many
SELECT s.id as schedule_id, s.type, s.start_time, s.end_time, u.full_name as dentist_name, r.name as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON s.id = a.schedule_id AND a.patient_id = sqlc.arg(patient_id)
WHERE s.start_time::date = sqlc.arg(date)::date
    AND s.slots_remaining > 0
    AND a.id IS NULL
ORDER BY s.start_time ASC;

-- name: ListExaminationSchedules :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       COUNT(a.id) AS appointment_count
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON s.id = a.schedule_id
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at ASC;

-- name: GetSchedule :one
SELECT s.id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       s.slots_remaining,
       s.created_at
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
WHERE s.id = sqlc.arg(schedule_id)
  AND s.type = sqlc.arg(type);

-- name: CreateSchedule :one
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id, slots_remaining)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateScheduleSlotsRemaining :exec
UPDATE schedules
SET slots_remaining = slots_remaining + $2
WHERE id = $1;

-- name: GetScheduleOverlap :many
SELECT s.id
FROM schedules s
WHERE s.room_id = $1
  AND s.start_time = $2
  AND s.end_time = $3;

-- name: ListPatientsByExaminationScheduleID :many
SELECT u.id,
       u.full_name,
       u.email,
       u.phone_number,
       u.date_of_birth,
       u.gender,
       u.role,
       sc.name as service_category
FROM users u
         JOIN appointments a ON u.id = a.patient_id
         JOIN schedules s ON a.schedule_id = s.id
         LEFT JOIN examination_appointment_detail ead ON a.id = ead.appointment_id
         JOIN service_categories sc ON ead.service_category_id = sc.id
WHERE s.id = sqlc.arg(schedule_id);

-- name: ListPatientsByTreatmentScheduleID :many
SELECT u.id,
       u.full_name,
       u.email,
       u.phone_number,
       u.date_of_birth,
       u.gender,
       u.role,
       services.name as service_name,
       tad.service_quantity
FROM users u
         JOIN appointments a ON u.id = a.patient_id
         JOIN schedules s ON a.schedule_id = s.id
         LEFT JOIN treatment_appointment_detail tad ON a.id = tad.appointment_id
         JOIN services ON tad.service_id = services.id
WHERE s.id = sqlc.arg(schedule_id);

-- name: ListExaminationSchedulesByDentistName :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       COUNT(a.id) AS appointment_count
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON s.id = a.schedule_id
WHERE u.full_name ILIKE '%' || sqlc.arg(dentist_name)::text || '%'
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at ASC;