// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"time"
)

type Querier interface {
	CreateAppointment(ctx context.Context, arg CreateAppointmentParams) error
	CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error)
	CreateDentist(ctx context.Context, arg CreateDentistParams) (User, error)
	CreateDentistDetail(ctx context.Context, arg CreateDentistDetailParams) (DentistDetail, error)
	CreateExaminationScheduleDetail(ctx context.Context, scheduleID int64) (ExaminationScheduleDetail, error)
	CreatePayment(ctx context.Context, name string) (Payment, error)
	CreateRoom(ctx context.Context, name string) (Room, error)
	CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error)
	CreateService(ctx context.Context, arg CreateServiceParams) (Service, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateSpecialty(ctx context.Context, name string) (Specialty, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetExaminationAppointmentDetails(ctx context.Context, arg GetExaminationAppointmentDetailsParams) (GetExaminationAppointmentDetailsRow, error)
	GetExaminationScheduleDetail(ctx context.Context, scheduleID int64) (GetExaminationScheduleDetailRow, error)
	GetPatient(ctx context.Context, id int64) (User, error)
	GetServiceCategoryBySlug(ctx context.Context, slug string) (ServiceCategory, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListExaminationBookings(ctx context.Context, patientID int64) ([]Booking, error)
	ListExaminationSchedulesByDate(ctx context.Context, date time.Time) ([]ListExaminationSchedulesByDateRow, error)
	ListPayments(ctx context.Context) ([]Payment, error)
	ListServiceCategories(ctx context.Context) ([]ServiceCategory, error)
	ListServicesOfOneCategory(ctx context.Context, slug string) ([]Service, error)
	UpdateExaminationScheduleSlotsRemaining(ctx context.Context, scheduleID int64) error
	UpdateServiceCategoryOfExaminationSchedule(ctx context.Context, arg UpdateServiceCategoryOfExaminationScheduleParams) error
}

var _ Querier = (*Queries)(nil)
