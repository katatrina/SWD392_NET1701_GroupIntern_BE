package db

import (
	"context"
)

type BookExaminationAppointmentParams struct {
	PatientID             int64
	PatientNote           string
	ExaminationScheduleID int64
	PaymentID             int64
}

func (store *SQLStore) BookExaminationAppointmentByPatientTx(ctx context.Context, arg BookExaminationAppointmentParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Create a new booking
		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			PatientID:   arg.PatientID,
			PatientNote: arg.PatientNote,
			PaymentID:   arg.PaymentID,
		})
		if err != nil {
			return err
		}

		// Create a new examination appointment
		err = q.CreateAppointment(ctx, CreateAppointmentParams{
			BookingID:  booking.ID,
			ScheduleID: arg.ExaminationScheduleID,
			PatientID:  arg.PatientID,
		})

		return err
	})

	return err
}
