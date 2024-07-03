// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: schedules.sql

package db

import (
	"context"
	"time"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id, slots_remaining)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, type, start_time, end_time, dentist_id, room_id, slots_remaining, created_at
`

type CreateScheduleParams struct {
	Type           string    `json:"type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	DentistID      int64     `json:"dentist_id"`
	RoomID         int64     `json:"room_id"`
	SlotsRemaining int64     `json:"slots_remaining"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
		arg.Type,
		arg.StartTime,
		arg.EndTime,
		arg.DentistID,
		arg.RoomID,
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
		&i.SlotsRemaining,
		&i.CreatedAt,
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
		&i.SlotsRemaining,
		&i.CreatedAt,
	)
	return i, err
}

const getScheduleOverlap = `-- name: GetScheduleOverlap :many
SELECT s.id
FROM schedules s
WHERE s.room_id = $1
  AND s.start_time = $2
  AND s.end_time = $3
`

type GetScheduleOverlapParams struct {
	RoomID    int64     `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
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
       COUNT(a.id) AS appointment_count
FROM schedules s
         JOIN users u ON s.dentist_id = u.id
         JOIN rooms r ON s.room_id = r.id
         LEFT JOIN appointments a ON s.id = a.schedule_id
GROUP BY s.id, u.full_name, r.name
ORDER BY s.created_at ASC
`

type ListExaminationSchedulesRow struct {
	ScheduleID       int64     `json:"schedule_id"`
	Type             string    `json:"type"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	DentistName      string    `json:"dentist_name"`
	RoomName         string    `json:"room_name"`
	AppointmentCount int64     `json:"appointment_count"`
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
			&i.AppointmentCount,
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
