-- name: CreateAppointment :one
INSERT INTO appointments (booking_id, schedule_id, patient_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetExaminationAppointmentDetails :one
SELECT b.id        as booking_id,
       b.type,
       b.status    as booking_status,
       b.payment_status,
       sc.name     as service_category,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       b.total_cost
FROM bookings b
         JOIN appointments a ON b.id = a.booking_id
         JOIN schedules s ON a.schedule_id = s.id
         JOIN examination_appointment_detail ead ON a.id = ead.appointment_id
         JOIN users u ON s.dentist_id = u.id
         JOIN dentist_detail dd ON u.id = dd.dentist_id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN service_categories sc ON ead.service_category_id = sc.id
WHERE b.id = sqlc.arg(booking_id)
  AND b.type = 'Examination'
  AND b.patient_id = sqlc.arg(patient_id);

-- name: CreateExaminationAppointmentDetail :one
INSERT INTO examination_appointment_detail (appointment_id, service_category_id)
VALUES ($1, $2) RETURNING *;

-- name: CreateTreatmentAppointmentDetail :one
INSERT INTO treatment_appointment_detail (appointment_id, service_id, service_quantity)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAppointmentByScheduleIDAndPatientID :one
SELECT *
FROM appointments
WHERE schedule_id = $1
  AND patient_id = $2;

-- name: UpdateAppointmentStatus :exec
UPDATE appointments
SET status = $2
WHERE id = $1;

-- name: GetAppointmentByBookingID :one
SELECT *
FROM appointments
WHERE booking_id = $1;

-- name: GetTreatmentAppointmentDetails :one
SELECT b.id          as booking_id,
       b.type,
       b.status      as booking_status,
       b.payment_status,
       services.name as service,
       tad.service_quantity,
       s.start_time,
       s.end_time,
       u.full_name   as dentist_name,
       r.name        as room_name,
       b.total_cost
FROM bookings b
         JOIN appointments a ON b.id = a.booking_id
         JOIN schedules s ON a.schedule_id = s.id
         JOIN treatment_appointment_detail tad ON a.id = tad.appointment_id
         JOIN users u ON s.dentist_id = u.id
         JOIN dentist_detail dd ON u.id = dd.dentist_id
         JOIN rooms r ON s.room_id = r.id
         JOIN services ON tad.service_id = services.id
WHERE b.id = sqlc.arg(booking_id)
  AND b.type = 'Treatment'
  AND b.patient_id = sqlc.arg(patient_id);