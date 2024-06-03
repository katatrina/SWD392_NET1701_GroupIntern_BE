package api

type createExaminationBookingRequest struct {
	CustomerID            int64  `json:"customer_id" binding:"required"`
	CustomerReason        string `json:"customer_reason"`
	PaymentID             int64  `json:"payment_id" binding:"required"`
	ExaminationScheduleID int64  `json:"examination_schedule_id" binding:"required"`
}

// createExaminationBooking creates a new examination booking
//
//	@Router		/bookings/examination [post]
//	@Summary	create a new examination booking
//	@Description
//	@Security	accessToken
//	@Tags		bookings
//	@Accept		json
//	@Produce	json
//	@Param		booking	body	createExaminationBookingRequest	true	"Create examination booking"
//	@Success	201
//	@Failure	400
//	@Failure	500
// func (server *Server) createExaminationBooking(ctx *gin.Context) {
// 	var req createExaminationBookingRequest
//
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
//
// 	authorizedPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
//
// 	customerID, err := strconv.ParseInt(authorizedPayload.Subject, 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	if req.CustomerID != customerID {
// 		err = errors.New("customer id doesn't match the authenticated user")
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}
//
// 	arg := db.BookExaminationAppointmentParams{
// 		CustomerID:            req.CustomerID,
// 		CustomerReason:        req.CustomerReason,
// 		PaymentID:             req.PaymentID,
// 		ExaminationScheduleID: req.ExaminationScheduleID,
// 	}
//
// 	err = server.store.BookExaminationAppointmentTx(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	ctx.JSON(http.StatusCreated, nil)
// }
