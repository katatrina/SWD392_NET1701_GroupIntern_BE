package api

import (
	"database/sql"
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/util"
)

// listDentists returns a list of dentists
//
//	@Router		/dentists [get]
//	@Summary	Lấy danh sách bác sĩ
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
	Sex         string          `json:"sex" binding:"required"`
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
//	@Success	201
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
	
	arg := db.CreateUserParams{
		FullName:       req.FullName,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Role:           "Dentist",
		HashedPassword: hashedPassword,
	}
	
	dentist, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	arg2 := db.CreateDentistDetailParams{
		DentistID:   dentist.ID,
		DateOfBirth: req.DateOfBirth.Time,
		Sex:         req.Sex,
		SpecialtyID: req.SpecialtyID,
	}
	
	_, err = server.store.CreateDentistDetail(ctx, arg2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
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
