package api

import (
	"errors"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern/db/sqlc"
)

var (
	ErrScheduleOverlap = errors.New("schedule overlaps with other schedules")
)

type createExaminationScheduleRequest struct {
	DentistID int64     `json:"dentist_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	RoomID    int64     `json:"room_id" binding:"required"`
}

// createExaminationSchedule creates a new examination schedule
//
//	@Router		/schedules/examination [post]
//	@Summary	Thêm lịch khám tổng quát
//	@Description
//	@Tags		schedules
//	@Accept		json
//	@Produce	json
//	@Param		request	body		createExaminationScheduleRequest	true	"Examination schedule information"
//	@Success	201		{object}	db.Schedule
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) createExaminationSchedule(ctx *gin.Context) {
	var req createExaminationScheduleRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// Check if the schedule overlaps with other schedules
	schedules, err := server.store.GetScheduleOverlap(ctx, db.GetScheduleOverlapParams{
		RoomID:    req.RoomID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if len(schedules) > 0 {
		ctx.JSON(http.StatusForbidden, errorResponse(ErrScheduleOverlap))
		return
	}
	
	schedule, err := server.store.CreateSchedule(ctx, db.CreateScheduleParams{
		Type:           "Examination",
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		DentistID:      req.DentistID,
		RoomID:         req.RoomID,
		SlotsRemaining: 3,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, schedule)
}

type listAvailableExaminationSchedulesByDateRequest struct {
	// Date time.Time `form:"date" time_format:"02/01/2006" binding:"required"`
	PatientID int64     `form:"patient_id" binding:"required"`
	Date      time.Time `form:"date" time_format:"2006-01-02" binding:"required"`
}

// listAvailableExaminationSchedulesByDate lists available examination schedules by date for a patient to book
//
//	@Router		/schedules/examination/available [get]
//	@Summary	Liệt kê tất cả lịch khám tổng quát còn trống trong một ngày cho bệnh nhân đặt lịch
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Param		date		query	string	true	"Date in the format YYYY-MM-DD"
//	@Param		patient_id	query	int		true	"Patient ID"
//	@Success	200			{array}	db.ListAvailableExaminationSchedulesByDateForPatientRow
//	@Failure	400
//	@Failure	500
func (server *Server) listAvailableExaminationSchedulesByDate(ctx *gin.Context) {
	var req listAvailableExaminationSchedulesByDateRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	schedules, err := server.store.ListAvailableExaminationSchedulesByDateForPatient(ctx, db.ListAvailableExaminationSchedulesByDateForPatientParams{
		PatientID: req.PatientID,
		Date:      req.Date,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, schedules)
}

// listExaminationSchedules lists all examination schedules
//
//	@Router		/schedules/examination [get]
//	@Summary	Liệt kê tất cả lịch khám tổng quát
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Success	200	{array}	db.ListExaminationSchedulesRow
//	@Failure	500
func (server *Server) listExaminationSchedules(ctx *gin.Context) {
	schedules, err := server.store.ListExaminationSchedules(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, schedules)
}
