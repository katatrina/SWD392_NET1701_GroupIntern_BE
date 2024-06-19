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

type updateServiceRequest struct {
	Name             *string `json:"name"`
	CategoryID       *int64  `json:"category_id"`
	Unit             *string `json:"unit"`
	Cost             *int64  `json:"cost"`
	WarrantyDuration *string `json:"warranty_duration"`
}

// updateService updates a service
//
//	@Router		/services/{id} [patch]
//	@Summary	Cập nhật thông tin của một dịch vụ
//	@Description
//	@Tags		services
//	@Accept		json
//	@Produce	json
//	@Param		request	body	updateServiceRequest	true	"Update service info"
//	@Param		id		path	int						true	"Service ID"
//	@Success	204
//	@Failure	400
//	@Failure	500
func (server *Server) updateService(ctx *gin.Context) {
	var req updateServiceRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	serviceID, err := server.getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	service, err := server.store.GetService(ctx, serviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if req.Name != nil {
		service.Name = *req.Name
	}
	if req.CategoryID != nil {
		service.CategoryID = *req.CategoryID
	}
	if req.Unit != nil {
		service.Unit = *req.Unit
	}
	if req.Cost != nil {
		service.Cost = *req.Cost
	}
	if req.WarrantyDuration != nil {
		service.WarrantyDuration = *req.WarrantyDuration
	}
	
	arg := db.UpdateServiceParams{
		ID:               service.ID,
		Name:             service.Name,
		CategoryID:       service.CategoryID,
		Unit:             service.Unit,
		Cost:             service.Cost,
		WarrantyDuration: service.WarrantyDuration,
	}
	
	err = server.store.UpdateService(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusNoContent, nil)
}

// getService returns a service
//
//	@Router		/services/{id} [get]
//	@Summary	Lấy thông tin của một dịch vụ
//	@Description
//	@Tags		services
//	@Produce	json
//	@Param		id	path		int	true	"Service ID"
//	@Success	200	{object}	db.Service
//	@Failure	400
//	@Failure	500
func (server *Server) getService(ctx *gin.Context) {
	serviceID, err := server.getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	service, err := server.store.GetService(ctx, serviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, service)
}
