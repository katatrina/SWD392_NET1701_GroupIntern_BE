package api

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
)

type listExaminationSchedulesRequest struct {
	// Date time.Time `form:"date" time_format:"02/01/2006" binding:"required"`
	Date time.Time `form:"date" time_format:"2006-01-02" binding:"required"`
}

// listExaminationSchedulesByDate lists examination schedules by date
//
//	@Router		/schedules/examination [get]
//	@Summary	Liệt kê tất cả lịch khám trong một ngày
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Param		date	query		string	true	"Date in the format YYYY-MM-DD"
//	@Success	200		{object}	[]db.ListExaminationSchedulesByDateRow
//	@Failure	400
//	@Failure	500
func (server *Server) listExaminationSchedulesByDate(ctx *gin.Context) {
	var req listExaminationSchedulesRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	schedules, err := server.store.ListExaminationSchedulesByDate(ctx, req.Date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, schedules)
}
