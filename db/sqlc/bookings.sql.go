// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: bookings.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createBooking = `-- name: CreateBooking :one
INSERT INTO bookings (patient_id, type, payment_status, payment_id, total_cost, appointment_date)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, patient_id, type, payment_status, payment_id, total_cost, appointment_date, status, created_at
`

type CreateBookingParams struct {
	PatientID       int64         `json:"patient_id"`
	Type            string        `json:"type"`
	PaymentStatus   string        `json:"payment_status"`
	PaymentID       sql.NullInt64 `json:"payment_id"`
	TotalCost       int64         `json:"total_cost"`
	AppointmentDate time.Time     `json:"appointment_date"`
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, createBooking,
		arg.PatientID,
		arg.Type,
		arg.PaymentStatus,
		arg.PaymentID,
		arg.TotalCost,
		arg.AppointmentDate,
	)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.PatientID,
		&i.Type,
		&i.PaymentStatus,
		&i.PaymentID,
		&i.TotalCost,
		&i.AppointmentDate,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listBookingsOfOnePatient = `-- name: ListBookingsOfOnePatient :many
SELECT id, patient_id, type, payment_status, payment_id, total_cost, appointment_date, status, created_at
FROM bookings
WHERE patient_id = $1
  AND type = $2
ORDER BY created_at DESC
`

type ListBookingsOfOnePatientParams struct {
	PatientID int64  `json:"patient_id"`
	Type      string `json:"type"`
}

func (q *Queries) ListBookingsOfOnePatient(ctx context.Context, arg ListBookingsOfOnePatientParams) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, listBookingsOfOnePatient, arg.PatientID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Booking{}
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.ID,
			&i.PatientID,
			&i.Type,
			&i.PaymentStatus,
			&i.PaymentID,
			&i.TotalCost,
			&i.AppointmentDate,
			&i.Status,
			&i.CreatedAt,
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

const updateBookingStatus = `-- name: UpdateBookingStatus :one
UPDATE bookings
SET status = $2
WHERE id = $1
RETURNING id, patient_id, type, payment_status, payment_id, total_cost, appointment_date, status, created_at
`

type UpdateBookingStatusParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateBookingStatus(ctx context.Context, arg UpdateBookingStatusParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, updateBookingStatus, arg.ID, arg.Status)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.PatientID,
		&i.Type,
		&i.PaymentStatus,
		&i.PaymentID,
		&i.TotalCost,
		&i.AppointmentDate,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
