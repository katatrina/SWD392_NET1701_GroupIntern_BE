package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// listAllPaymentMethods returns a list of all payment methods
//
//	@Router		/payment-methods [get]
//	@Summary	liệt kê tất cả phương thức thanh toán
//	@Produce	json
//	@Description
//	@Tags		payments
//	@Success	200 {object} []db.Payment
//	@Failure	500
func (server *Server) listAllPaymentMethods(ctx *gin.Context) {
	payments, err := server.store.ListAllPaymentMethods(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
