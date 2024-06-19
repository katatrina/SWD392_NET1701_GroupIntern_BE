package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
)

type createServiceRequest struct {
	Name             string `json:"name" binding:"required"`
	CategoryID       int64  `json:"category_id" binding:"required"`
	Unit             string `json:"unit" binding:"required"`
	Cost             int64  `json:"cost" binding:"required"`
	WarrantyDuration string `json:"warranty_duration" binding:"required"`
}

// createService creates a new service
//
//	@Router		/services [post]
//	@Summary	Tạo mới dịch vụ
//	@Description
//	@Tags		services
//	@Accept		json
//	@Produce	json
//	@Param		request	body		createServiceRequest	true	"Create service info"
//	@Success	201		{object}	db.Service
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) createService(ctx *gin.Context) {
	var req createServiceRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	arg := db.CreateServiceParams{
		Name:             req.Name,
		CategoryID:       req.CategoryID,
		Unit:             req.Unit,
		Cost:             req.Cost,
		WarrantyDuration: req.WarrantyDuration,
	}
	
	_, err := server.store.CreateService(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
}
