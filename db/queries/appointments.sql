-- name: CreateAppointment :exec
INSERT INTO appointments (booking_id, schedule_id, patient_id)
VALUES ($1, $2, $3);

-- name: ListExaminationAppointments :many
SELECT schedules.start_time, bookings.id as booking_id, service_categories.price as fee, bookings.status as status
FROM bookings
         JOIN appointments ON bookings.id = appointments.booking_id
         JOIN schedules ON appointments.schedule_id = schedules.id
         JOIN examination_schedule_detail ON schedules.id = examination_schedule_detail.schedule_id
         JOIN service_categories ON examination_schedule_detail.service_category_id = service_categories.id
WHERE bookings.patient_id = $1
ORDER BY schedules.start_time DESC
LIMIT $2
OFFSET $3;