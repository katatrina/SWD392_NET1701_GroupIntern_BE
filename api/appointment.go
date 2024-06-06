package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/token"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

type createExaminationAppointmentByPatientRequest struct {
	PatientNote           string `json:"patient_note"`
	ExaminationScheduleID int64  `json:"examination_schedule_id" binding:"required"`
	PaymentID             int64  `json:"payment_id" binding:"required"`
}

// createExaminationAppointmentByPatient creates a new examination appointment for a patient
//
//	@Router		/patients/me/appointments/examination [post]
//	@Summary	Cho phép bệnh nhân đặt lịch khám
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

	authorizedPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	patientID, err := strconv.ParseInt(authorizedPayload.Subject, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.BookExaminationAppointmentParams{
		PatientID:             patientID,
		PatientNote:           req.PatientNote,
		ExaminationScheduleID: req.ExaminationScheduleID,
		PaymentID:             req.PaymentID,
	}

	err = server.store.BookExaminationAppointmentByPatientTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

type listExaminationAppointmentsByPatientRequest struct {
	PageID   int32 `form:"page_id" binding:"required"`
	PageSize int32 `form:"page_size" binding:"required"`
}

// listExaminationAppointmentsByPatient returns a list of examination appointments of a patient
//
//	@Router		/patients/me/appointments/examination [get]
//	@Summary	Lấy danh sách lịch khám của bệnh nhân
//	@Produce	json
//	@Description
//	@Security	accessToken
//	@Param		page_id		query	int	true	"Page ID"
//	@Param		page_size	query	int	true	"Page Size"
//	@Tags		appointments
//	@Success	200	{object}	[]db.ListExaminationAppointmentsRow
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) listExaminationAppointmentsByPatient(ctx *gin.Context) {
	var req listExaminationAppointmentsByPatientRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authorizedPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	patientID, err := strconv.ParseInt(authorizedPayload.Subject, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.ListExaminationAppointmentsParams{
		PatientID: patientID,
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
	}

	appointments, err := server.store.ListExaminationAppointments(ctx, arg)
	// TODO: Handle no record error gracefully
	if len(appointments) == 0 {
		ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}
