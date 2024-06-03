package db

import (
	"context"
)

type BookExaminationAppointmentParams struct {
	CustomerID            int64
	CustomerNote          string
	ExaminationScheduleID int64
	PaymentID             int64
}

func (store *SQLStore) BookExaminationAppointmentTx(ctx context.Context, arg BookExaminationAppointmentParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Create a new booking
		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			CustomerID:   arg.CustomerID,
			CustomerNote: arg.CustomerNote,
			PaymentID:    arg.PaymentID,
		})
		if err != nil {
			return err
		}

		// Create a new examination appointment
		err = q.CreateAppointment(ctx, CreateAppointmentParams{
			BookingID:  booking.ID,
			ScheduleID: arg.ExaminationScheduleID,
			CustomerID: arg.CustomerID,
		})

		return err
	})

	return err
}
