package db

import (
	"context"
	"database/sql"
)

type BookExaminationAppointmentParams struct {
	PatientID             int64
	ExaminationScheduleID int64
	ServiceCategoryID     int64
}

func (store *SQLStore) BookExaminationAppointmentByPatientTx(ctx context.Context, arg BookExaminationAppointmentParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Get examination schedule detail
		schedule, err := q.GetExaminationScheduleDetail(ctx, arg.ExaminationScheduleID)
		if err != nil {
			return err
		}
		
		// Create a new booking
		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			PatientID:     arg.PatientID,
			Type:          "Examination",
			PaymentStatus: "Không cần thanh toán",
			PaymentID: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			TotalCost:       0,
			AppointmentDate: schedule.StartTime,
		})
		if err != nil {
			return err
		}
		
		// Create a new examination appointment
		err = q.CreateAppointment(ctx, CreateAppointmentParams{
			BookingID:  booking.ID,
			ScheduleID: schedule.ID,
			PatientID:  arg.PatientID,
		})
		if err != nil {
			return err
		}
		
		// Update service category ID
		if arg.ServiceCategoryID > 0 {
			err = q.UpdateServiceCategoryOfExaminationSchedule(ctx, UpdateServiceCategoryOfExaminationScheduleParams{
				ScheduleID: schedule.ID,
				ServiceCategoryID: sql.NullInt64{
					Int64: arg.ServiceCategoryID,
					Valid: true,
				},
			})
			if err != nil {
				return err
			}
		}
		
		// Update slots remaining
		err = q.UpdateExaminationScheduleSlotsRemaining(ctx, schedule.ID)
		if err != nil {
			return err
		}
		
		return err
	})
	
	return err
}
