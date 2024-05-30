package db

import (
	"context"
	"database/sql"
)

type BookExaminationAppointmentParams struct {
	CustomerID            int64  `json:"customer_id"`
	CustomerReason        string `json:"customer_reason"`
	PaymentID             int64  `json:"payment_id"`
	ExaminationScheduleID int64  `json:"examination_schedule_id"`
}

func (store *SQLStore) BookExaminationAppointmentTx(ctx context.Context, arg BookExaminationAppointmentParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Create a new examination booking
		booking, err := q.CreateExaminationBooking(ctx, CreateExaminationBookingParams{
			CustomerID:     arg.CustomerID,
			CustomerReason: arg.CustomerReason,
			PaymentID:      arg.PaymentID,
		})
		if err != nil {
			return err
		}

		// Update examination schedule
		err = q.UpdateExaminationSchedule(ctx, UpdateExaminationScheduleParams{
			ID: arg.ExaminationScheduleID,
			BookingID: sql.NullInt64{
				Int64: booking.ID,
				Valid: true,
			},
			CustomerID: sql.NullInt64{
				Int64: arg.CustomerID,
				Valid: true,
			},
		})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
