package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
)

type listExaminationSchedulesRequest struct {
	Date              time.Time `form:"date" time_format:"2006-01-02" binding:"required"`
	ServiceCategoryID int64     `form:"service_category_id" binding:"required"`
}

// listExaminationSchedulesByDateAndServiceCategory lists examination schedules by date and service category
//
//	@Router		/schedules/examination [get]
//	@Summary	list examination schedules by date and service category
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Param		date				query	string	true	"Date in the format YYYY-MM-DD"
//	@Param		service_category_id	query	int		true	"Service Category ID"
//	@Success	200
//	@Failure	400
//	@Failure	500
func (server *Server) listExaminationSchedulesByDateAndServiceCategory(ctx *gin.Context) {
	var req listExaminationSchedulesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListExaminationSchedulesByDateAndServiceCategoryParams{
		Date:              req.Date,
		ServiceCategoryID: req.ServiceCategoryID,
	}

	schedules, err := server.store.ListExaminationSchedulesByDateAndServiceCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedules)
}
