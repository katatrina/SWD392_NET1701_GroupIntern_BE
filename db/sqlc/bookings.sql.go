// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: bookings.sql

package db

import (
	"context"
	"time"
)

const createExaminationBooking = `-- name: CreateExaminationBooking :one
INSERT INTO bookings (patient_id, patient_note, payment_id, total_cost, appointment_date, type)
VALUES ($1, $2, $3, $4, $5, 'Examination') RETURNING id, patient_id, patient_note, type, payment_status, payment_id, total_cost, appointment_date, status, created_at
`

type CreateExaminationBookingParams struct {
	PatientID       int64     `json:"patient_id"`
	PatientNote     string    `json:"patient_note"`
	PaymentID       int64     `json:"payment_id"`
	TotalCost       int64     `json:"total_cost"`
	AppointmentDate time.Time `json:"appointment_date"`
}

func (q *Queries) CreateExaminationBooking(ctx context.Context, arg CreateExaminationBookingParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, createExaminationBooking,
		arg.PatientID,
		arg.PatientNote,
		arg.PaymentID,
		arg.TotalCost,
		arg.AppointmentDate,
	)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.PatientID,
		&i.PatientNote,
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

const listExaminationBookings = `-- name: ListExaminationBookings :many
SELECT id, patient_id, patient_note, type, payment_status, payment_id, total_cost, appointment_date, status, created_at
FROM bookings
WHERE patient_id = $1
  AND type = 'Examination'
`

func (q *Queries) ListExaminationBookings(ctx context.Context, patientID int64) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, listExaminationBookings, patientID)
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
			&i.PatientNote,
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
