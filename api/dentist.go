package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/util"
)

type dentistResponse struct {
	ID          int64           `json:"id"`
	FullName    string          `json:"full_name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	DateOfBirth util.CustomDate `json:"date_of_birth"`
	Gender      string          `json:"gender"`
	Specialty   string          `json:"specialty"`
}

// listDentists returns a list of dentists
//
//	@Router		/dentists [get]
//	@Summary	Lấy danh sách nha sĩ
//	@Produce	json
//	@Param		q	query	string	false	"Search query by name"
//	@Description
//	@Tags		dentists
//	@Success	200	{array}	db.ListDentistsRow
//	@Failure	404
//	@Failure	500
func (server *Server) listDentists(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	if searchQuery == "" {
		services, err := server.store.ListDentists(ctx)
		switch {
		case err != nil:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		case len(services) == 0:
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
		default:
			ctx.JSON(http.StatusOK, services)
		}
		
		return
	}
	
	services, err := server.store.ListDentistsByName(ctx, searchQuery)
	switch {
	case err != nil:
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	case len(services) == 0:
		ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
	default:
		ctx.JSON(http.StatusOK, services)
	}
}

type createDentistRequest struct {
	FullName    string          `json:"full_name" binding:"required"`
	Email       string          `json:"email" binding:"required,email"`
	PhoneNumber string          `json:"phone_number" binding:"required"`
	DateOfBirth util.CustomDate `json:"date" binding:"required"`
	Gender      string          `json:"gender" binding:"required"`
	SpecialtyID int64           `json:"specialty_id" binding:"required"`
	Password    string          `json:"password" binding:"required"`
}

// createDentist creates a new dentist
//
//	@Router		/dentists [post]
//	@Summary	Tạo mới nha sĩ
//	@Produce	json
//	@Accept		json
//	@Param		request	body	createDentistRequest	true	"Create dentist info"
//	@Description
//	@Tags		dentists
//	@Success	201	{object}	db.CreateDentistAccountResult
//	@Failure	400
//	@Failure	500
func (server *Server) createDentist(ctx *gin.Context) {
	var req createDentistRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	hashedPassword, err := util.GenerateHashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	arg := db.CreateDentistAccountParams{
		FullName:       req.FullName,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		SpecialtyID:    req.SpecialtyID,
		HashedPassword: hashedPassword,
	}
	
	result, err := server.store.CreateDentistAccountTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, result)
}

// getDentist returns a dentist by ID
//
//	@Router		/dentists/{id} [get]
//	@Summary	Lấy thông tin nha sĩ
//	@Produce	json
//	@Param		id	path	int	true	"Dentist ID"
//	@Description
//	@Tags		dentists
//	@Success	200	{object}	db.GetDentistRow
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) getDentist(ctx *gin.Context) {
	dentistID, err := server.getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	dentist, err := server.store.GetDentist(ctx, dentistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, dentist)
}

type updateDentistRequest struct {
	FullName    *string          `json:"full_name"`
	Email       *string          `json:"email"`
	PhoneNumber *string          `json:"phone_number"`
	DateOfBirth *util.CustomDate `json:"date_of_birth"`
	Gender      *string          `json:"gender"`
	SpecialtyID *int64           `json:"specialty_id"`
}

// updateDentistProfile updates a dentist's profile
//
//	@Router		/dentists/profile [patch]
//	@Summary	Cập nhật thông tin nha sĩ
//	@Produce	json
//	@Accept		json
//	@Security	accessToken
//	@Param		request	body	updateDentistRequest	true	"Update dentist info"
//	@Description
//	@Tags		dentists
//	@Success	200	{object}	db.UpdateDentistProfileResult
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) updateDentistProfile(ctx *gin.Context) {
	dentistID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	dentist, err := server.store.GetDentist(ctx, dentistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	var req updateDentistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	if req.FullName != nil {
		dentist.FullName = *req.FullName
	}
	if req.Email != nil {
		dentist.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		dentist.PhoneNumber = *req.PhoneNumber
	}
	if req.DateOfBirth != nil {
		dentist.DateOfBirth = time.Time(*req.DateOfBirth)
	}
	if req.Gender != nil {
		dentist.Gender = *req.Gender
	}
	if req.SpecialtyID != nil {
		dentist.SpecialtyID = *req.SpecialtyID
	}
	
	arg := db.UpdateDentistProfileParams{
		DentistID:   dentist.ID,
		FullName:    dentist.FullName,
		Email:       dentist.Email,
		PhoneNumber: dentist.PhoneNumber,
		DateOfBirth: dentist.DateOfBirth,
		Gender:      dentist.Gender,
		SpecialtyID: dentist.SpecialtyID,
	}
	
	result, err := server.store.UpdateDentistProfileTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, result)
}

// getDentistProfile returns the profile of the authorized dentist
//
//	@Router		/dentists/profile [get]
//	@Summary	Xem thông tin cá nhân nha sĩ
//	@Produce	json
//	@Security	accessToken
//	@Description
//	@Tags		dentists
//	@Success	200	{object}	db.GetDentistRow
//	@Failure	400
//	@Failure	404
//	@Failure	500
func (server *Server) getDentistProfile(ctx *gin.Context) {
	dentistID, err := server.getAuthorizedUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	dentist, err := server.store.GetDentist(ctx, dentistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, dentist)
}
