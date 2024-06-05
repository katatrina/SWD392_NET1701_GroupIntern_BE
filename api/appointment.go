package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/token"
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
//	@Param		request	body	createExaminationAppointmentByPatientRequest	true "Examination Appointment Request"
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

	customerID, err := strconv.ParseInt(authorizedPayload.Subject, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.BookExaminationAppointmentParams{
		PatientID:             customerID,
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

type listExaminationAppointmentsRequest struct {
	CustomerID int64 `uri:"id" binding:"required"`
}

func (server *Server) listExaminationAppointmentsByPatient(ctx *gin.Context) {

}
