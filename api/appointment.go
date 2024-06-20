package api

import (
	"database/sql"
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

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
//	@Tags		appointments
//	@Accept		json
//	@Produce	json
//	@Param		request	body	createExaminationAppointmentByPatientRequest	true	"Examination Appointment Request"
//	@Success	201
//	@Failure	400
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
//	@Tags		appointments
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
//	@Tags		appointments
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
