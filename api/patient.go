package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/db/sqlc"
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/util"
)

type createPatientRequest struct {
	Password    string          `json:"password" binding:"required"`
	FullName    string          `json:"full_name" binding:"required"`
	Email       string          `json:"email" binding:"required"`
	PhoneNumber string          `json:"phone_number" binding:"required"`
	DateOfBirth util.CustomDate `json:"date_of_birth" binding:"required"`
	Gender      string          `json:"gender" binding:"required"`
}

// createPatient creates a new patient account
//
//	@Router		/patients [post]
//	@Summary	Tạo tài khoản bệnh nhân
//	@Description
//	@Tags		patients
//	@Accept		json
//	@Produce	json
//	@Param		request	body		createPatientRequest	true	"Create patient info"
//	@Success	201		{object}	db.User					"Patient account"
//	@Failure	400
//	@Failure	403	{object}	util.MapErrors	"Unique validation errors"
//	@Failure	500
func (server *Server) createPatient(ctx *gin.Context) {
	var req createPatientRequest
	
	// Parse the JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	errs := make(util.MapErrors)
	
	// Check if the email is existed
	emailExisted, err := server.store.IsEmailExists(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if emailExisted {
		errs.Add("email_error", ErrEmailAlreadyExist.Error())
	}
	
	// Check if the phone number is existed
	phoneNumberExisted, err := server.store.IsPhoneNumberExists(ctx, req.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if phoneNumberExisted {
		errs.Add("phone_number_error", ErrPhoneNumberAlreadyExist.Error())
	}
	
	// Return the error response if there are any validation errors
	if len(errs) > 0 {
		ctx.JSON(http.StatusForbidden, errs)
		return
	}
	
	// Generate hashed password
	hashedPassword, err := util.GenerateHashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	arg := db.CreateUserParams{
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Role:           "Patient",
		DateOfBirth:    time.Time(req.DateOfBirth),
		Gender:         req.Gender,
	}
	
	// Create a new patient account
	patient, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, patient)
}

// getPatient returns the information of a patient
//
//	@Router		/patients/{id} [get]
//	@Summary	Lấy thông tin cá nhân của bệnh nhân
//	@Description
//	@Tags		patients
//	@Param		id	path	int	true	"Patient ID"
//	@Produce	json
//	@Success	200	{object}	userInfo
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) getPatient(ctx *gin.Context) {
	patientID, err := server.getLastIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	patient, err := server.store.GetPatient(ctx, patientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	rsp := userInfo{
		ID:          patient.ID,
		FullName:    patient.FullName,
		Email:       patient.Email,
		PhoneNumber: patient.PhoneNumber,
		Role:        patient.Role,
		DateOfBirth: util.CustomDate(patient.DateOfBirth),
		Gender:      patient.Gender,
		CreatedAt:   patient.CreatedAt,
	}
	
	ctx.JSON(http.StatusOK, rsp)
}

type createExaminationAppointmentByPatientRequest struct {
	ExaminationScheduleID int64 `json:"examination_schedule_id" binding:"required"`
	ServiceCategoryID     int64 `json:"service_category_id"`
}

// createExaminationAppointmentByPatient creates a new examination appointment for a patient
//
//	@Router		/patients/appointments/examination [post]
//	@Summary	Cho phép bệnh nhân đặt lịch khám tổng quát
//	@Description
//	@Security	accessToken
//	@Tags		patients
//	@Accept		json
//	@Produce	json
//	@Param		request	body	createExaminationAppointmentByPatientRequest	true	"Examination Appointment Request"
//	@Success	201
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) createExaminationAppointmentByPatient(ctx *gin.Context) {
	var req createExaminationAppointmentByPatientRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	patientID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// TODO: Authorize the patient more strictly
	
	// Get schedule
	schedule, err := server.store.GetSchedule(ctx, db.GetScheduleParams{
		ScheduleID: req.ExaminationScheduleID,
		Type:       "Examination",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	// Check if the schedule is booked by the patient before
	_, err = server.store.GetAppointmentByScheduleIDAndPatientID(ctx, db.GetAppointmentByScheduleIDAndPatientIDParams{
		ScheduleID: req.ExaminationScheduleID,
		PatientID:  patientID,
	})
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusForbidden, errorResponse(db.ErrScheduleBookedByPatientBefore))
		return
	}
	
	// Check if the schedule is full
	if schedule.SlotsRemaining == 0 {
		ctx.JSON(http.StatusForbidden, errorResponse(db.ErrScheduleFullSlot))
		return
	}
	
	arg := db.BookExaminationScheduleParams{
		PatientID:         patientID,
		Schedule:          schedule,
		ServiceCategoryID: req.ServiceCategoryID,
	}
	
	err = server.store.BookExaminationAppointmentByPatientTx(ctx, arg)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrScheduleBookedByPatientBefore):
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		case errors.Is(err, db.ErrScheduleFullSlot):
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
}

// listExaminationAppointmentsByPatient returns all examination bookings of a patient
//
//	@Router		/patients/appointments/examination [get]
//	@Summary	Cho phép bệnh nhân xem lịch sử tất cả lịch khám tổng quát của mình
//	@Produce	json
//	@Description
//	@Security	accessToken
//	@Tags		patients
//	@Success	200	{object}	[]db.Booking	"List of examination bookings"
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) listExaminationAppointmentsByPatient(ctx *gin.Context) {
	patientID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	arg := db.ListBookingsOfOnePatientParams{
		PatientID: patientID,
		Type:      "Examination",
	}
	
	bookings, err := server.store.ListBookingsOfOnePatient(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, bookings)
}

type getExaminationAppointmentDetailsRequest struct {
	ExaminationAppointmentID int64 `uri:"id" binding:"required"`
}

// getExaminationAppointmentByPatient returns the details of an examination appointment
//
//	@Router		/patients/appointments/examination/{id} [get]
//	@Summary	Cho phép bệnh nhân xem chi tiết một lịch khám tổng quát của mình
//	@Description
//	@Tags		patients
//	@Produce	json
//	@Security	accessToken
//	@Param		id	path		int	true	"Examination Appointment ID"
//	@Success	200	{object}	db.GetExaminationAppointmentDetailsRow
//	@Failure	400
//	@Failure	500
func (server *Server) getExaminationAppointmentByPatient(ctx *gin.Context) {
	var req getExaminationAppointmentDetailsRequest
	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	patientID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	arg := db.GetExaminationAppointmentDetailsParams{
		PatientID: patientID,
		BookingID: req.ExaminationAppointmentID,
	}
	
	details, err := server.store.GetExaminationAppointmentDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, details)
}

// cmancelExaminationAppointmentByPatient cancels an examination appointment by a patient
//
//	@Router		/appointments/examination/{id}/cancel [patch]
//	@Summary	Cho phép bệnh nhân hủy lịch khám
//	@Description
//	@Tags		appointments
//	@Param		id	path	int	true	"Examination Booking ID"
//	@Success	204
//	@Failure	400
//	@Failure	500
func (server *Server) cancelExaminationAppointmentByPatient(ctx *gin.Context) {
	bookingID, err := server.getMiddleIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// patientID, err := server.getAuthorizedUserID(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	
	arg := db.CancelExaminationAppointmentByPatientParams{
		// PatientID: patientID,
		BookingID: bookingID,
	}
	
	err = server.store.CancelExaminationAppointmentByPatientTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusNoContent, nil)
}
