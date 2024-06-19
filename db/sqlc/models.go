// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Appointment struct {
	ID         int64     `json:"id"`
	BookingID  int64     `json:"booking_id"`
	ScheduleID int64     `json:"schedule_id"`
	PatientID  int64     `json:"patient_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type Booking struct {
	ID              int64         `json:"id"`
	PatientID       int64         `json:"patient_id"`
	Type            string        `json:"type"`
	PaymentStatus   string        `json:"payment_status"`
	PaymentID       sql.NullInt64 `json:"payment_id"`
	TotalCost       int64         `json:"total_cost"`
	AppointmentDate time.Time     `json:"appointment_date"`
	Status          string        `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
}

type DentistDetail struct {
	DentistID   int64     `json:"dentist_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Sex         string    `json:"sex"`
	SpecialtyID int64     `json:"specialty_id"`
}

type ExaminationScheduleDetail struct {
	ScheduleID        int64         `json:"schedule_id"`
	ServiceCategoryID sql.NullInt64 `json:"service_category_id"`
	SlotsRemaining    int64         `json:"slots_remaining"`
	CreatedAt         time.Time     `json:"created_at"`
}

type Payment struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Room struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Schedule struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	DentistID int64     `json:"dentist_id"`
	RoomID    int64     `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Service struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	CategoryID       int64     `json:"category_id"`
	Unit             string    `json:"unit"`
	Cost             int64     `json:"cost"`
	WarrantyDuration string    `json:"warranty_duration"`
	CreatedAt        time.Time `json:"created_at"`
}

type ServiceCategory struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	IconUrl     string    `json:"icon_url"`
	BannerUrl   string    `json:"banner_url"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"created_at"`
}

type Specialty struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type TreatmentScheduleDetail struct {
	ScheduleID      int64     `json:"schedule_id"`
	ServiceID       int64     `json:"service_id"`
	ServiceQuantity int64     `json:"service_quantity"`
	SlotRemains     int64     `json:"slot_remains"`
	CreatedAt       time.Time `json:"created_at"`
}

type User struct {
	ID             int64     `json:"id"`
	FullName       string    `json:"full_name"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
}
