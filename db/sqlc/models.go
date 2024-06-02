// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Booking struct {
	ID             int64     `json:"id"`
	Type           string    `json:"type"`
	CustomerID     int64     `json:"customer_id"`
	CustomerReason string    `json:"customer_reason"`
	PaymentStatus  string    `json:"payment_status"`
	PaymentID      int64     `json:"payment_id"`
	IsCancelled    bool      `json:"is_cancelled"`
	CreatedAt      time.Time `json:"created_at"`
}

type DentistDetail struct {
	DentistID   int64     `json:"dentist_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Sex         string    `json:"sex"`
	SpecialtyID int64     `json:"specialty_id"`
}

type ExaminationSchedule struct {
	ID                int64         `json:"id"`
	BookingID         sql.NullInt64 `json:"booking_id"`
	StartTime         time.Time     `json:"start_time"`
	EndTime           time.Time     `json:"end_time"`
	CustomerID        sql.NullInt64 `json:"customer_id"`
	DentistID         int64         `json:"dentist_id"`
	ServiceCategoryID int64         `json:"service_category_id"`
	RoomID            int64         `json:"room_id"`
	Slot              int64         `json:"slot"`
	Status            string        `json:"status"`
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

type Service struct {
	ID               int64     `json:"id"`
	CategoryID       int64     `json:"category_id"`
	Unit             string    `json:"unit"`
	Price            int64     `json:"price"`
	WarrantyDuration int64     `json:"warranty_duration"`
	CreatedAt        time.Time `json:"created_at"`
}

type ServiceCategory struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	Slug      string    `json:"slug"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type Specialty struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type TreatmentSchedule struct {
	ID              int64     `json:"id"`
	BookingID       int64     `json:"booking_id"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	CustomerID      int64     `json:"customer_id"`
	DentistID       int64     `json:"dentist_id"`
	ServiceID       int64     `json:"service_id"`
	ServiceQuantity int64     `json:"service_quantity"`
	RoomID          int64     `json:"room_id"`
	Slot            int64     `json:"slot"`
	Status          string    `json:"status"`
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