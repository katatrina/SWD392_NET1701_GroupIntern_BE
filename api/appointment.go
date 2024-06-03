package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/token"
)

type createExaminationAppointmentRequest struct {
	CustomerID            int64  `json:"customer_id" binding:"required"`
	CustomerNote          string `json:"customer_note"`
	ExaminationScheduleID int64  `json:"examination_schedule_id" binding:"required"`
	PaymentID             int64  `json:"payment_id" binding:"required"`
}

// createExaminationAppointment creates a new examination appointment
//
//	@Router		/appointments/examination [post]
//	@Summary	create a new examination appointment
//	@Description
//	@Security	accessToken
//	@Tags		appointments
//	@Accept		json
//	@Produce	json
//	@Param		booking	body	createExaminationAppointmentRequest	true "Examination Appointment Request"
//	@Success	201
//	@Failure	400
//	@Failure	500
func (server *Server) createExaminationAppointment(ctx *gin.Context) {
	var req createExaminationAppointmentRequest

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

	if req.CustomerID != customerID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrMismatchedUser))
		return
	}

	arg := db.BookExaminationAppointmentParams{
		CustomerID:            req.CustomerID,
		CustomerNote:          req.CustomerNote,
		ExaminationScheduleID: req.ExaminationScheduleID,
		PaymentID:             req.PaymentID,
	}

	err = server.store.BookExaminationAppointmentTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
