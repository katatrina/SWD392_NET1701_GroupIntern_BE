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
		// Get examination schedule detail
		schedule, err := q.GetExaminationScheduleDetail(ctx, arg.ExaminationScheduleID)
		if err != nil {
			return err
		}
		
		// Create a new booking
		booking, err := q.CreateExaminationBooking(ctx, CreateExaminationBookingParams{
			PatientID:       arg.PatientID,
			PatientNote:     arg.PatientNote,
			PaymentID:       arg.PaymentID,
			TotalCost:       schedule.ServiceCategoryCost,
			AppointmentDate: schedule.StartTime,
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
		
		// Update slots remaining
		err = q.UpdateExaminationScheduleSlotsRemaining(ctx, schedule.ScheduleID)
		if err != nil {
			return err
		}
		
		return err
	})
	
	return err
}
