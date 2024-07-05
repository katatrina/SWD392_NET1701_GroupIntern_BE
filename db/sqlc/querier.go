// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (Appointment, error)
	CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error)
	CreateDentistDetail(ctx context.Context, arg CreateDentistDetailParams) (DentistDetail, error)
	CreateExaminationAppointmentDetail(ctx context.Context, arg CreateExaminationAppointmentDetailParams) (ExaminationAppointmentDetail, error)
	CreatePayment(ctx context.Context, name string) (Payment, error)
	CreateRoom(ctx context.Context, name string) (Room, error)
	CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error)
	CreateService(ctx context.Context, arg CreateServiceParams) (Service, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateSpecialty(ctx context.Context, name string) (Specialty, error)
	CreateTreatmentAppointmentDetail(ctx context.Context, arg CreateTreatmentAppointmentDetailParams) (TreatmentAppointmentDetail, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteDentist(ctx context.Context, id int64) error
	DeleteService(ctx context.Context, id int64) error
	DeleteServiceCategory(ctx context.Context, id int64) error
	GetAppointmentByBookingID(ctx context.Context, bookingID int64) (Appointment, error)
	GetAppointmentByScheduleIDAndPatientID(ctx context.Context, arg GetAppointmentByScheduleIDAndPatientIDParams) (Appointment, error)
	GetDentist(ctx context.Context, id int64) (GetDentistRow, error)
	GetExaminationAppointmentDetails(ctx context.Context, arg GetExaminationAppointmentDetailsParams) (GetExaminationAppointmentDetailsRow, error)
	GetPatient(ctx context.Context, id int64) (User, error)
	GetSchedule(ctx context.Context, arg GetScheduleParams) (GetScheduleRow, error)
	GetScheduleOverlap(ctx context.Context, arg GetScheduleOverlapParams) ([]int64, error)
	GetService(ctx context.Context, id int64) (Service, error)
	GetServiceCategoryByID(ctx context.Context, id int64) (GetServiceCategoryByIDRow, error)
	GetServiceCategoryBySlug(ctx context.Context, slug string) (ServiceCategory, error)
	GetSpecialty(ctx context.Context, id int64) (Specialty, error)
	GetUserByEmailForLogin(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	ListAvailableExaminationSchedulesByDateForPatient(ctx context.Context, arg ListAvailableExaminationSchedulesByDateForPatientParams) ([]ListAvailableExaminationSchedulesByDateForPatientRow, error)
	ListBookingsOfOnePatient(ctx context.Context, arg ListBookingsOfOnePatientParams) ([]Booking, error)
	ListDentists(ctx context.Context) ([]ListDentistsRow, error)
	ListDentistsByName(ctx context.Context, name string) ([]ListDentistsByNameRow, error)
	ListExaminationSchedules(ctx context.Context) ([]ListExaminationSchedulesRow, error)
	ListPayments(ctx context.Context) ([]Payment, error)
	ListRooms(ctx context.Context) ([]Room, error)
	ListServiceCategories(ctx context.Context) ([]ServiceCategory, error)
	ListServiceCategoriesByName(ctx context.Context, name string) ([]ServiceCategory, error)
	ListServicesByCategory(ctx context.Context, slug string) ([]Service, error)
	ListServicesByNameAndCategory(ctx context.Context, arg ListServicesByNameAndCategoryParams) ([]Service, error)
	ListSpecialties(ctx context.Context) ([]Specialty, error)
	UpdateAppointmentStatus(ctx context.Context, arg UpdateAppointmentStatusParams) error
	UpdateBookingStatus(ctx context.Context, arg UpdateBookingStatusParams) (Booking, error)
	UpdateDentistDetail(ctx context.Context, arg UpdateDentistDetailParams) (DentistDetail, error)
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) error
	UpdateScheduleSlotsRemaining(ctx context.Context, arg UpdateScheduleSlotsRemainingParams) error
	UpdateService(ctx context.Context, arg UpdateServiceParams) error
	UpdateServiceCategory(ctx context.Context, arg UpdateServiceCategoryParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
