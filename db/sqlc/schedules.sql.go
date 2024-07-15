// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: schedules.sql

package db

import (
	"context"
	"time"

	util "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/util"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id, max_patients, slots_remaining)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, type, start_time, end_time, dentist_id, room_id, max_patients, slots_remaining, created_at
`

type CreateScheduleParams struct {
	Type           string    `json:"type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	DentistID      int64     `json:"dentist_id"`
	RoomID         int64     `json:"room_id"`
	MaxPatients    int64     `json:"max_patients"`
	SlotsRemaining int64     `json:"slots_remaining"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
		arg.Type,
		arg.StartTime,
		arg.EndTime,
		arg.DentistID,
		arg.RoomID,
		arg.MaxPatients,
		arg.SlotsRemaining,
	)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.StartTime,
		&i.EndTime,
		&i.DentistID,
		&i.RoomID,
		&i.MaxPatients,
		&i.SlotsRemaining,
		&i.CreatedAt,
	)
	return i, err
}

const getPatientByTreatmentScheduleID = `-- name: GetPatientByTreatmentScheduleID :one
SELECT u.id,
       u.full_name,
       u.email,
       u.phone_number,
       u.date_of_birth,
       u.gender,
       u.role,
       services.name as service_name,
       services.cost as service_cost,
       tad.service_quantity,
       b.total_cost
FROM users u
         JOIN appointments a ON u.id = a.patient_id
         JOIN schedules s ON a.schedule_id = s.id
         JOIN treatment_appointment_detail tad ON a.id = tad.appointment_id
         JOIN services ON tad.service_id = services.id
         JOIN bookings b ON a.booking_id = b.id
WHERE s.id = $1
  AND s.type = 'Treatment'
`

type GetPatientByTreatmentScheduleIDRow struct {
	ID              int64     `json:"id"`
	FullName        string    `json:"full_name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Gender          string    `json:"gender"`
	Role            string    `json:"role"`
	ServiceName     string    `json:"service_name"`
	ServiceCost     int64     `json:"service_cost"`
	ServiceQuantity int64     `json:"service_quantity"`
	TotalCost       int64     `json:"total_cost"`
}

func (q *Queries) GetPatientByTreatmentScheduleID(ctx context.Context, scheduleID int64) (GetPatientByTreatmentScheduleIDRow, error) {
	row := q.db.QueryRowContext(ctx, getPatientByTreatmentScheduleID, scheduleID)
	var i GetPatientByTreatmentScheduleIDRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.Gender,
		&i.Role,
		&i.ServiceName,
		&i.ServiceCost,
		&i.ServiceQuantity,
		&i.TotalCost,
	)
	return i, err
}

const getSchedule = `-- name: GetSchedule :one
SELECT s.id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       s.max_patients,
       s.slots_remaining,
       s.created_at
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
WHERE s.id = $1
  AND s.type = $2
`

type GetScheduleParams struct {
	ScheduleID int64  `json:"schedule_id"`
	Type       string `json:"type"`
}

type GetScheduleRow struct {
	ID             int64     `json:"id"`
	Type           string    `json:"type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	DentistName    string    `json:"dentist_name"`
	RoomName       string    `json:"room_name"`
	MaxPatients    int64     `json:"max_patients"`
	SlotsRemaining int64     `json:"slots_remaining"`
	CreatedAt      time.Time `json:"created_at"`
}

func (q *Queries) GetSchedule(ctx context.Context, arg GetScheduleParams) (GetScheduleRow, error) {
	row := q.db.QueryRowContext(ctx, getSchedule, arg.ScheduleID, arg.Type)
	var i GetScheduleRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.StartTime,
		&i.EndTime,
		&i.DentistName,
		&i.RoomName,
		&i.MaxPatients,
		&i.SlotsRemaining,
		&i.CreatedAt,
	)
	return i, err
}

const getScheduleOverlap = `-- name: GetScheduleOverlap :many
SELECT s.id
FROM schedules s
WHERE s.room_id = $1
  AND (s.start_time, s.end_time) OVERLAPS ($2, $3)
`

type GetScheduleOverlapParams struct {
	RoomID    int64       `json:"room_id"`
	StartTime interface{} `json:"start_time"`
	EndTime   interface{} `json:"end_time"`
}

func (q *Queries) GetScheduleOverlap(ctx context.Context, arg GetScheduleOverlapParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getScheduleOverlap, arg.RoomID, arg.StartTime, arg.EndTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAvailableExaminationSchedulesByDateForPatient = `-- name: ListAvailableExaminationSchedulesByDateForPatient :many
SELECT s.id as schedule_id, s.type, s.start_time, s.end_time, u.full_name as dentist_name, r.name as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON s.id = a.schedule_id AND a.patient_id = $1
WHERE s.start_time::date = $2::date
    AND s.slots_remaining > 0
    AND a.id IS NULL
ORDER BY s.start_time ASC
`

type ListAvailableExaminationSchedulesByDateForPatientParams struct {
	PatientID int64     `json:"patient_id"`
	Date      time.Time `json:"date"`
}

type ListAvailableExaminationSchedulesByDateForPatientRow struct {
	ScheduleID  int64     `json:"schedule_id"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	DentistName string    `json:"dentist_name"`
	RoomName    string    `json:"room_name"`
}

func (q *Queries) ListAvailableExaminationSchedulesByDateForPatient(ctx context.Context, arg ListAvailableExaminationSchedulesByDateForPatientParams) ([]ListAvailableExaminationSchedulesByDateForPatientRow, error) {
	rows, err := q.db.QueryContext(ctx, listAvailableExaminationSchedulesByDateForPatient, arg.PatientID, arg.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAvailableExaminationSchedulesByDateForPatientRow{}
	for rows.Next() {
		var i ListAvailableExaminationSchedulesByDateForPatientRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExaminationSchedules = `-- name: ListExaminationSchedules :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       s.max_patients,
       COUNT(a.id) as current_patients
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON (s.id = a.schedule_id AND a.status <> 'Đã hủy')
WHERE s.type = 'Examination'
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at DESC
`

type ListExaminationSchedulesRow struct {
	ScheduleID      int64     `json:"schedule_id"`
	Type            string    `json:"type"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	DentistName     string    `json:"dentist_name"`
	RoomName        string    `json:"room_name"`
	MaxPatients     int64     `json:"max_patients"`
	CurrentPatients int64     `json:"current_patients"`
}

func (q *Queries) ListExaminationSchedules(ctx context.Context) ([]ListExaminationSchedulesRow, error) {
	rows, err := q.db.QueryContext(ctx, listExaminationSchedules)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExaminationSchedulesRow{}
	for rows.Next() {
		var i ListExaminationSchedulesRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
			&i.MaxPatients,
			&i.CurrentPatients,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExaminationSchedulesByDentistID = `-- name: ListExaminationSchedulesByDentistID :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       s.max_patients,
       COUNT(a.id) as current_patients
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON (s.id = a.schedule_id AND a.status <> 'Đã hủy')
WHERE u.id = $1
  AND s.type = 'Examination'
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at DESC
`

type ListExaminationSchedulesByDentistIDRow struct {
	ScheduleID      int64     `json:"schedule_id"`
	Type            string    `json:"type"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	DentistName     string    `json:"dentist_name"`
	RoomName        string    `json:"room_name"`
	MaxPatients     int64     `json:"max_patients"`
	CurrentPatients int64     `json:"current_patients"`
}

func (q *Queries) ListExaminationSchedulesByDentistID(ctx context.Context, dentistID int64) ([]ListExaminationSchedulesByDentistIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listExaminationSchedulesByDentistID, dentistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExaminationSchedulesByDentistIDRow{}
	for rows.Next() {
		var i ListExaminationSchedulesByDentistIDRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
			&i.MaxPatients,
			&i.CurrentPatients,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExaminationSchedulesByDentistName = `-- name: ListExaminationSchedulesByDentistName :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name,
       s.max_patients,
       COUNT(a.id) as current_patients
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON (s.id = a.schedule_id AND a.status <> 'Đã hủy')
WHERE u.full_name ILIKE '%' || $1::text || '%'
AND s.type = 'Examination'
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at DESC
`

type ListExaminationSchedulesByDentistNameRow struct {
	ScheduleID      int64     `json:"schedule_id"`
	Type            string    `json:"type"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	DentistName     string    `json:"dentist_name"`
	RoomName        string    `json:"room_name"`
	MaxPatients     int64     `json:"max_patients"`
	CurrentPatients int64     `json:"current_patients"`
}

func (q *Queries) ListExaminationSchedulesByDentistName(ctx context.Context, dentistName string) ([]ListExaminationSchedulesByDentistNameRow, error) {
	rows, err := q.db.QueryContext(ctx, listExaminationSchedulesByDentistName, dentistName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExaminationSchedulesByDentistNameRow{}
	for rows.Next() {
		var i ListExaminationSchedulesByDentistNameRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
			&i.MaxPatients,
			&i.CurrentPatients,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPatientsByExaminationScheduleID = `-- name: ListPatientsByExaminationScheduleID :many
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
         LEFT JOIN service_categories sc ON ead.service_category_id = sc.id
WHERE s.id = $1
  AND s.type = 'Examination'
  AND a.status <> 'Đã hủy'
`

type ListPatientsByExaminationScheduleIDRow struct {
	ID              int64               `json:"id"`
	FullName        string              `json:"full_name"`
	Email           string              `json:"email"`
	PhoneNumber     string              `json:"phone_number"`
	DateOfBirth     time.Time           `json:"date_of_birth"`
	Gender          string              `json:"gender"`
	Role            string              `json:"role"`
	ServiceCategory util.JSONNullString `json:"service_category"`
}

func (q *Queries) ListPatientsByExaminationScheduleID(ctx context.Context, scheduleID int64) ([]ListPatientsByExaminationScheduleIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listPatientsByExaminationScheduleID, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPatientsByExaminationScheduleIDRow{}
	for rows.Next() {
		var i ListPatientsByExaminationScheduleIDRow
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.PhoneNumber,
			&i.DateOfBirth,
			&i.Gender,
			&i.Role,
			&i.ServiceCategory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTreatmentSchedules = `-- name: ListTreatmentSchedules :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
WHERE s.type = 'Treatment'
ORDER BY s.created_at DESC
`

type ListTreatmentSchedulesRow struct {
	ScheduleID  int64     `json:"schedule_id"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	DentistName string    `json:"dentist_name"`
	RoomName    string    `json:"room_name"`
}

func (q *Queries) ListTreatmentSchedules(ctx context.Context) ([]ListTreatmentSchedulesRow, error) {
	rows, err := q.db.QueryContext(ctx, listTreatmentSchedules)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTreatmentSchedulesRow{}
	for rows.Next() {
		var i ListTreatmentSchedulesRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTreatmentSchedulesByDentistID = `-- name: ListTreatmentSchedulesByDentistID :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
WHERE u.id = $1
  AND s.type = 'Treatment'
ORDER BY s.created_at DESC
`

type ListTreatmentSchedulesByDentistIDRow struct {
	ScheduleID  int64     `json:"schedule_id"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	DentistName string    `json:"dentist_name"`
	RoomName    string    `json:"room_name"`
}

func (q *Queries) ListTreatmentSchedulesByDentistID(ctx context.Context, dentistID int64) ([]ListTreatmentSchedulesByDentistIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listTreatmentSchedulesByDentistID, dentistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTreatmentSchedulesByDentistIDRow{}
	for rows.Next() {
		var i ListTreatmentSchedulesByDentistIDRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTreatmentSchedulesByDentistName = `-- name: ListTreatmentSchedulesByDentistName :many
SELECT s.id        as schedule_id,
       s.type,
       s.start_time,
       s.end_time,
       u.full_name as dentist_name,
       r.name      as room_name
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
WHERE u.full_name ILIKE '%' || $1::text || '%'
AND s.type = 'Treatment'
ORDER BY s.created_at DESC
`

type ListTreatmentSchedulesByDentistNameRow struct {
	ScheduleID  int64     `json:"schedule_id"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	DentistName string    `json:"dentist_name"`
	RoomName    string    `json:"room_name"`
}

func (q *Queries) ListTreatmentSchedulesByDentistName(ctx context.Context, dentistName string) ([]ListTreatmentSchedulesByDentistNameRow, error) {
	rows, err := q.db.QueryContext(ctx, listTreatmentSchedulesByDentistName, dentistName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTreatmentSchedulesByDentistNameRow{}
	for rows.Next() {
		var i ListTreatmentSchedulesByDentistNameRow
		if err := rows.Scan(
			&i.ScheduleID,
			&i.Type,
			&i.StartTime,
			&i.EndTime,
			&i.DentistName,
			&i.RoomName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateScheduleSlotsRemaining = `-- name: UpdateScheduleSlotsRemaining :exec
UPDATE schedules
SET slots_remaining = slots_remaining + $2
WHERE id = $1
`

type UpdateScheduleSlotsRemainingParams struct {
	ID             int64 `json:"id"`
	SlotsRemaining int64 `json:"slots_remaining"`
}

func (q *Queries) UpdateScheduleSlotsRemaining(ctx context.Context, arg UpdateScheduleSlotsRemainingParams) error {
	_, err := q.db.ExecContext(ctx, updateScheduleSlotsRemaining, arg.ID, arg.SlotsRemaining)
	return err
}
