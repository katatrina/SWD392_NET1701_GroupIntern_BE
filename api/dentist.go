package api

import (
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
//	@Description
//	@Tags		dentists
//	@Success	200	{array}	db.ListDentistsRow
//	@Failure	500
func (server *Server) listDentists(ctx *gin.Context) {
	dentists, err := server.store.ListDentists(ctx)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, dentists)
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
//	@Summary	Tạo mới bác sĩ
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
