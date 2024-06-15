-- name: CreateAppointment :exec
INSERT INTO appointments (booking_id, schedule_id, patient_id)
VALUES ($1, $2, $3);

-- name: GetExaminationAppointmentDetails :one
SELECT b.id as booking_id, b.status as booking_status, b.payment_status, b.patient_note, s.start_time, s.end_time, u.full_name as dentist_name, specialties.name as dentist_specialty, r.name as room_name, b.total_cost
FROM bookings b
         JOIN appointments a ON b.id = a.booking_id
         JOIN schedules s ON a.schedule_id = s.id
         JOIN examination_schedule_detail sd ON s.id = sd.schedule_id
         JOIN users u ON s.dentist_id = u.id
         JOIN dentist_detail dd ON u.id = dd.dentist_id
         JOIN specialties ON dd.specialty_id = specialties.id
         JOIN rooms r ON s.room_id = r.id
         JOIN service_categories sc ON sd.service_category_id = sc.id
WHERE b.id = sqlc.arg(booking_id) AND b.patient_id = sqlc.arg(patient_id);
