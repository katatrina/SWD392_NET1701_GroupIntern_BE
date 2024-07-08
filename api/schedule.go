package api

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern/db/sqlc"
)

var ()

type createExaminationScheduleRequest struct {
	DentistID int64     `json:"dentist_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	RoomID    int64     `json:"room_id" binding:"required"`
}

// createExaminationSchedule creates a new examination schedule
//
//	@Router		/schedules/examination [post]
//	@Summary	Tạo lịch khám tổng quát bởi Admin
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

// listAvailableExaminationSchedulesByDateForPatient lists available examination schedules by date for a patient to book
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
func (server *Server) listAvailableExaminationSchedulesByDateForPatient(ctx *gin.Context) {
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

type createTreatmentScheduleRequest struct {
	DentistID       int64     `json:"dentist_id" binding:"required"`
	PatientID       int64     `json:"patient_id" binding:"required"`
	StartTime       time.Time `json:"start_time" binding:"required"`
	EndTime         time.Time `json:"end_time" binding:"required"`
	RoomID          int64     `json:"room_id" binding:"required"`
	ServiceID       int64     `json:"service_id" binding:"required"`
	ServiceQuantity int64     `json:"service_quantity" binding:"required"`
	PaymentID       int64     `json:"payment_id" binding:"required"`
}

// createTreatmentSchedule creates a new treatment schedule
//
//	@Router		/schedules/treatment [post]
//	@Summary	Tạo lịch điều trị bởi nha sĩ
//	@Description
//	@Tags		schedules
//	@Security	accessToken
//	@Accept		json
//	@Produce	json
//	@Param		request	body	createTreatmentScheduleRequest	true	"Treatment schedule information"
//	@Success	201
//	@Failure	400
//	@Failure	401
//	@Failure	403
//	@Failure	500
func (server *Server) createTreatmentSchedule(ctx *gin.Context) {
	dentistID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	
	var req createTreatmentScheduleRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// Check if the dentist is the same as the dentist in the request
	if dentistID != req.DentistID {
		ctx.JSON(http.StatusForbidden, errorResponse(ErrMisMatchedUserID))
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
	
	err = server.store.BookTreatmentAppointmentByDentistTx(ctx, db.BookTreatmentAppointmentByDentistTxParams{
		DentistID:       req.DentistID,
		PatientID:       req.PatientID,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		RoomID:          req.RoomID,
		ServiceID:       req.ServiceID,
		ServiceQuantity: req.ServiceQuantity,
		PaymentID:       req.PaymentID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
}

// listPatientsByExaminationSchedule lists patients by examination schedule
//
//	@Router		/schedules/examination/{id}/patients [get]
//	@Summary	Liệt kê tất cả bệnh nhân của một lịch khám tổng quát
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Param		id	path	int	true	"Examination Schedule ID"
//	@Success	200	{array}	db.ListPatientsByExaminationScheduleIDRow
//	@Failure	400
//	@Failure	500
func (server *Server) listPatientsByExaminationSchedule(ctx *gin.Context) {
	scheduleID, err := server.getMiddleIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	patients, err := server.store.ListPatientsByExaminationScheduleID(ctx, scheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, patients)
}

// listPatientsByExaminationSchedule lists patients by examination schedule
//
//	@Router		/schedules/treatment/{id}/patients [get]
//	@Summary	Liệt kê tất cả bệnh nhân của một lịch điều trị
//	@Description
//	@Tags		schedules
//	@Produce	json
//	@Param		id	path	int	true	"Treatment Schedule ID"
//	@Success	200	{array}	db.ListPatientsByTreatmentScheduleIDRow
//	@Failure	400
//	@Failure	500
func (server *Server) listPatientsByTreatmentSchedule(ctx *gin.Context) {
	scheduleID, err := server.getMiddleIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	patients, err := server.store.ListPatientsByTreatmentScheduleID(ctx, scheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, patients)
}
