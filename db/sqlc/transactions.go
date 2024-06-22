package db

import (
	"context"
	"database/sql"
	"time"
	
	"github.com/katatrina/SWD392/internal/util"
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
		})
		if err != nil {
			return err
		}
		result.DentistID = dentist.ID
		result.FullName = dentist.FullName
		result.Email = dentist.Email
		result.PhoneNumber = dentist.PhoneNumber
		
		// Create dentist detail
		dentistDetail, err := q.CreateDentistDetail(ctx, CreateDentistDetailParams{
			DentistID:   dentist.ID,
			DateOfBirth: time.Time(arg.DateOfBirth),
			Gender:      arg.Gender,
			SpecialtyID: arg.SpecialtyID,
		})
		if err != nil {
			return err
		}
		result.DateOfBirth = util.CustomDate(dentistDetail.DateOfBirth)
		result.Gender = dentistDetail.Gender
		
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
			Role:        "Dentist",
		})
		if err != nil {
			return err
		}
		result.DentistID = dentist.ID
		result.FullName = dentist.FullName
		result.Email = dentist.Email
		result.PhoneNumber = dentist.PhoneNumber
		
		// Update dentist detail
		dentistDetail, err := q.UpdateDentistDetail(ctx, UpdateDentistDetailParams{
			DentistID:   arg.DentistID,
			DateOfBirth: arg.DateOfBirth,
			Gender:      arg.Gender,
			SpecialtyID: arg.SpecialtyID,
		})
		if err != nil {
			return err
		}
		result.DateOfBirth = dentistDetail.DateOfBirth
		result.Gender = dentistDetail.Gender
		
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
