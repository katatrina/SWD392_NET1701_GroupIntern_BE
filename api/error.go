package api

import (
	"errors"
)

var (
	ErrEmailNotFound           = errors.New("email not found")
	ErrEmailAlreadyExist       = errors.New("email already exists")
	ErrPhoneNumberAlreadyExist = errors.New("phone number already exists")
	ErrPasswordIncorrect       = errors.New("password is incorrect")
	ErrMisMatchedUserID        = errors.New("the provided user ID does not match the authorized user ID")
	
	ErrNoRecordFound   = errors.New("no record found")
	ErrScheduleOverlap = errors.New("schedule overlaps with other schedules")
	
	ErrEmptySearchQuery = errors.New("search query is empty")
)
