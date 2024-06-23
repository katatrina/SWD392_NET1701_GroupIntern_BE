package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/util"
	"github.com/lib/pq"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

type createPatientRequest struct {
	Password    string `json:"password" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// createPatient creates a new patient account
//
//	@Router		/patients [post]
//	@Summary	Tạo tài khoản bệnh nhân
//	@Description
//	@Tags		patients
//	@Accept		json
//	@Produce	json
//	@Param		request	body	createPatientRequest	true	"Create patient info"
//	@Success	201
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) createPatient(ctx *gin.Context) {
	var req createPatientRequest
	
	// Parse the JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
	}
	
	// Create a new customer
	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch {
			case pqErr.Code.Name() == "unique_violation":
				err = fmt.Errorf("%s", pqErr.Detail)
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			default:
				err = fmt.Errorf("unexpected error occured: %s", pqErr.Detail)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
}

// getPatientProfile returns the information of a patient
//
//	@Router		/patients/profile [get]
//	@Summary	Lấy thông tin bệnh nhân
//	@Description
//	@Tags		patients
//	@Produce	json
//	@Security	accessToken
//	@Success	200	{object}	userInfo
//	@Failure	400
//	@Failure	403
//	@Failure	404
//	@Failure	500
func (server *Server) getPatientProfile(ctx *gin.Context) {
	patientID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
	
	arg := db.BookExaminationAppointmentParams{
		PatientID:             patientID,
		ExaminationScheduleID: req.ExaminationScheduleID,
		ServiceCategoryID:     req.ServiceCategoryID,
	}
	
	err = server.store.BookExaminationAppointmentByPatientTx(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrExaminationScheduleFull) {
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
//	@Summary	Lấy tất cả danh sách lịch khám của bệnh nhân
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
	
	arg := db.ListBookingsParams{
		PatientID: patientID,
		Type:      "Examination",
	}
	
	bookings, err := server.store.ListBookings(ctx, arg)
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
//	@Summary	Lấy thông tin chi tiết của một lịch khám
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
