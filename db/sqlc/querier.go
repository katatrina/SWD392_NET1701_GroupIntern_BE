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
	CreateDentistDetail(ctx context.Context, arg CreateDentistDetailParams) (DentistDetail, error)
	CreateExaminationScheduleDetail(ctx context.Context, scheduleID int64) (ExaminationScheduleDetail, error)
	CreatePayment(ctx context.Context, name string) (Payment, error)
	CreateRoom(ctx context.Context, name string) (Room, error)
	CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error)
	CreateService(ctx context.Context, arg CreateServiceParams) (Service, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateSpecialty(ctx context.Context, name string) (Specialty, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteService(ctx context.Context, id int64) error
	DeleteServiceCategory(ctx context.Context, id int64) error
	GetExaminationAppointmentDetails(ctx context.Context, arg GetExaminationAppointmentDetailsParams) (GetExaminationAppointmentDetailsRow, error)
	GetExaminationScheduleDetail(ctx context.Context, scheduleID int64) (GetExaminationScheduleDetailRow, error)
	GetPatient(ctx context.Context, id int64) (User, error)
	GetService(ctx context.Context, id int64) (Service, error)
	GetServiceCategoryByID(ctx context.Context, id int64) (ServiceCategory, error)
	GetServiceCategoryBySlug(ctx context.Context, slug string) (ServiceCategory, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListBookings(ctx context.Context, arg ListBookingsParams) ([]Booking, error)
	ListExaminationSchedulesByDate(ctx context.Context, date time.Time) ([]ListExaminationSchedulesByDateRow, error)
	ListPayments(ctx context.Context) ([]Payment, error)
	ListServiceCategories(ctx context.Context) ([]ServiceCategory, error)
	ListServices(ctx context.Context) ([]Service, error)
	ListServicesByCategory(ctx context.Context, slug string) ([]Service, error)
	ListServicesByName(ctx context.Context, name string) ([]Service, error)
	UpdateExaminationScheduleSlotsRemaining(ctx context.Context, scheduleID int64) error
	UpdateService(ctx context.Context, arg UpdateServiceParams) error
	UpdateServiceCategory(ctx context.Context, arg UpdateServiceCategoryParams) error
	UpdateServiceCategoryOfExaminationSchedule(ctx context.Context, arg UpdateServiceCategoryOfExaminationScheduleParams) error
}

var _ Querier = (*Queries)(nil)
