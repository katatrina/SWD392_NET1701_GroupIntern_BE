package db

import (
	"context"
	"database/sql"
	"errors"
	"time"
	
	"github.com/katatrina/SWD392_NET1701_GroupIntern/internal/util"
)

var (
	ErrScheduleFullSlot              = errors.New("schedule is full slot")
	ErrScheduleBookedByPatientBefore = errors.New("schedule is booked by the patient before")
)

type BookExaminationScheduleParams struct {
	PatientID         int64
	Schedule          GetScheduleRow
	ServiceCategoryID int64
}

func (store *SQLStore) BookExaminationAppointmentByPatientTx(ctx context.Context, arg BookExaminationScheduleParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
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
			AppointmentDate: arg.Schedule.StartTime,
		})
		if err != nil {
			return err
		}
		
		// Create a new examination appointment
		appointment, err := q.CreateAppointment(ctx, CreateAppointmentParams{
			BookingID:  booking.ID,
			ScheduleID: arg.Schedule.ID,
			PatientID:  arg.PatientID,
		})
		if err != nil {
			return err
		}
		
		// Update service category ID of the examination appointment
		if arg.ServiceCategoryID > 0 {
			_, err := q.CreateExaminationAppointmentDetail(ctx, CreateExaminationAppointmentDetailParams{
				AppointmentID: appointment.ID,
				ServiceCategoryID: util.JSONNullInt64{
					NullInt64: sql.NullInt64{
						Int64: arg.ServiceCategoryID,
						Valid: true,
					}},
			})
			if err != nil {
				return err
			}
		} else {
			_, err := q.CreateExaminationAppointmentDetail(ctx, CreateExaminationAppointmentDetailParams{
				AppointmentID: appointment.ID,
				ServiceCategoryID: util.JSONNullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					}},
			})
			if err != nil {
				return err
			}
		}
		
		// Update slots remaining
		err = q.UpdateScheduleSlotsRemaining(ctx, UpdateScheduleSlotsRemainingParams{
			ID:             arg.Schedule.ID,
			SlotsRemaining: -1,
		})
		if err != nil {
			return err
		}
		
		return err
	})
	
	return err
}

type CreateDentistAccountParams struct {
	FullName       string          `json:"full_name"`
	Email          string          `json:"email"`
	PhoneNumber    string          `json:"phone_number"`
	DateOfBirth    util.CustomDate `json:"date"`
	Gender         string          `json:"gender"`
	SpecialtyID    int64           `json:"specialty_id"`
	HashedPassword string          `json:"hashed_password"`
}

type CreateDentistAccountResult struct {
	DentistID   int64           `json:"dentist_id"`
	FullName    string          `json:"full_name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Gender      string          `json:"gender"`
	DateOfBirth util.CustomDate `json:"date_of_birth"`
	Specialty   string          `json:"specialty"`
}

func (store *SQLStore) CreateDentistAccountTx(ctx context.Context, arg CreateDentistAccountParams) (CreateDentistAccountResult, error) {
	var result CreateDentistAccountResult
	
	err := store.execTx(ctx, func(q *Queries) error {
		// Create a new dentist account
		dentist, err := q.CreateUser(ctx, CreateUserParams{
			FullName:       arg.FullName,
			Email:          arg.Email,
			PhoneNumber:    arg.PhoneNumber,
			Role:           "Dentist",
			HashedPassword: arg.HashedPassword,
			DateOfBirth:    time.Time(arg.DateOfBirth),
			Gender:         arg.Gender,
		})
		if err != nil {
			return err
		}
		result.DentistID = dentist.ID
		result.FullName = dentist.FullName
		result.Email = dentist.Email
		result.PhoneNumber = dentist.PhoneNumber
		result.DateOfBirth = util.CustomDate(dentist.DateOfBirth)
		result.Gender = dentist.Gender
		
		// Create dentist detail
		_, err = q.CreateDentistDetail(ctx, CreateDentistDetailParams{
			DentistID: dentist.ID,
			
			SpecialtyID: arg.SpecialtyID,
		})
		if err != nil {
			return err
		}
		
		// Get specialty name
		specialty, err := q.GetSpecialty(ctx, arg.SpecialtyID)
		if err != nil {
			return err
		}
		result.Specialty = specialty.Name
		
		return nil
	})
	
	return result, err
}

type UpdateDentistProfileParams struct {
	DentistID   int64     `json:"dentist_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	SpecialtyID int64     `json:"specialty_id"`
}

type UpdateDentistProfileResult struct {
	DentistID   int64     `json:"dentist_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Specialty   string    `json:"specialty"`
}

func (store *SQLStore) UpdateDentistProfileTx(ctx context.Context, arg UpdateDentistProfileParams) (UpdateDentistProfileResult, error) {
	var result UpdateDentistProfileResult
	
	err := store.execTx(ctx, func(q *Queries) error {
		// Update dentist account
		dentist, err := q.UpdateUser(ctx, UpdateUserParams{
			ID:          arg.DentistID,
			FullName:    arg.FullName,
			Email:       arg.Email,
			PhoneNumber: arg.PhoneNumber,
			DateOfBirth: arg.DateOfBirth,
			Gender:      arg.Gender,
		})
		if err != nil {
			return err
		}
		result.DentistID = dentist.ID
		result.FullName = dentist.FullName
		result.Email = dentist.Email
		result.PhoneNumber = dentist.PhoneNumber
		result.DateOfBirth = dentist.DateOfBirth
		result.Gender = dentist.Gender
		
		// Update dentist detail
		_, err = q.UpdateDentistDetail(ctx, UpdateDentistDetailParams{
			DentistID:   arg.DentistID,
			SpecialtyID: arg.SpecialtyID,
		})
		if err != nil {
			return err
		}
		
		// Get specialty name
		specialty, err := q.GetSpecialty(ctx, arg.SpecialtyID)
		if err != nil {
			return err
		}
		result.Specialty = specialty.Name
		
		return nil
	})
	
	return result, err
}

type BookTreatmentAppointmentByDentistTxParams struct {
	DentistID       int64     `json:"dentist_id"`
	PatientID       int64     `json:"patient_id"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	RoomID          int64     `json:"room_id"`
	ServiceID       int64     `json:"service_id"`
	ServiceQuantity int64     `json:"service_quantity"`
	PaymentID       int64     `json:"payment_id"`
}

func (store *SQLStore) BookTreatmentAppointmentByDentistTx(ctx context.Context, arg BookTreatmentAppointmentByDentistTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Create a new treatment schedule
		schedule, err := q.CreateSchedule(ctx, CreateScheduleParams{
			Type:           "Treatment",
			StartTime:      arg.StartTime,
			EndTime:        arg.EndTime,
			DentistID:      arg.DentistID,
			RoomID:         arg.RoomID,
			SlotsRemaining: 1,
		})
		if err != nil {
			return err
		}
		
		// Get service
		service, err := q.GetService(ctx, arg.ServiceID)
		if err != nil {
			return err
		}
		
		// Create a new booking
		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			PatientID:     arg.PatientID,
			Type:          "Treatment",
			PaymentStatus: "Chưa thanh toán",
			PaymentID: sql.NullInt64{
				Int64: arg.PaymentID,
				Valid: true,
			},
			TotalCost:       service.Cost * arg.ServiceQuantity,
			AppointmentDate: arg.StartTime,
		})
		if err != nil {
			return err
		}
		
		// Create a new treatment appointment
		appointment, err := q.CreateAppointment(ctx, CreateAppointmentParams{
			BookingID:  booking.ID,
			ScheduleID: schedule.ID,
			PatientID:  arg.PatientID,
		})
		if err != nil {
			return err
		}
		
		// Create a new treatment appointment detail
		_, err = q.CreateTreatmentAppointmentDetail(ctx, CreateTreatmentAppointmentDetailParams{
			AppointmentID:   appointment.ID,
			ServiceID:       arg.ServiceID,
			ServiceQuantity: arg.ServiceQuantity,
		})
		if err != nil {
			return err
		}
		
		return nil
	})
	
	return err
}

type CancelExaminationAppointmentByPatientParams struct {
	PatientID int64 `json:"patient_id"`
	BookingID int64 `json:"booking_id"`
}

func (store *SQLStore) CancelExaminationAppointmentByPatientTx(ctx context.Context, arg CancelExaminationAppointmentByPatientParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Update booking status
		booking, err := q.UpdateBookingStatus(ctx, UpdateBookingStatusParams{
			ID:     arg.BookingID,
			Status: "Đã hủy",
		})
		if err != nil {
			return err
		}
		
		// Get appointment
		appointment, err := q.GetAppointmentByBookingID(ctx, booking.ID)
		if err != nil {
			return err
		}
		
		// Update appointment status
		err = q.UpdateAppointmentStatus(ctx, UpdateAppointmentStatusParams{
			ID:     appointment.ID,
			Status: "Đã hủy",
		})
		
		// Update slots remaining
		err = q.UpdateScheduleSlotsRemaining(ctx, UpdateScheduleSlotsRemainingParams{
			ID:             appointment.ScheduleID,
			SlotsRemaining: +1,
		})
		if err != nil {
			return err
		}
		return nil
	})
	
	return err
}
