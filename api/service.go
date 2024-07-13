package api

import (
	"errors"
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/db/sqlc"
	"github.com/lib/pq"
)

type createServiceRequest struct {
	Name             string `json:"name" binding:"required"`
	CategoryID       int64  `json:"category_id" binding:"required"`
	Unit             string `json:"unit" binding:"required"`
	Cost             int64  `json:"cost" binding:"required"`
	WarrantyDuration string `json:"warranty_duration"`
}

// createService creates a new service
//
//	@Router		/services [post]
//	@Summary	Thêm một dịch vụ
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
		Currency:         "VND",
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
//	@Router		/services/{id} [put]
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
	
	serviceID, err := server.getLastIDParam(ctx)
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
	serviceID, err := server.getLastIDParam(ctx)
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

// deleteService deletes a service
//
//	@Router		/services/{id} [delete]
//	@Summary	Xóa một dịch vụ
//	@Description
//	@Param		id	path	int	true	"Service ID"
//	@Tags		services
//	@Success	204
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) deleteService(ctx *gin.Context) {
	serviceID, err := server.getLastIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	err = server.store.DeleteService(ctx, serviceID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch {
			case pqErr.Code.Name() == "foreign_key_violation":
				err = fmt.Errorf("%s", pqErr.Detail)
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			default:
				err = fmt.Errorf("unexpected error occured: %s", pqErr.Detail)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusNoContent, nil)
}

type listServicesRequest struct {
	CategorySlug string `form:"category" binding:"required"`
	SearchQuery  string `form:"q"`
}

// listServices returns a list of services
//
//	@Router		/services [get]
//	@Summary	Lấy danh sách dịch vụ của một loại hình
//	@Description
//	@Tags		services
//	@Produce	json
//	@Param		category	query	string	true	"Filter services by category slug"
//	@Param		q			query	string	false	"Search query by service name"
//	@Success	200			{array}	db.Service
//	@Failure	404
//	@Failure	500
func (server *Server) listServices(ctx *gin.Context) {
	var req listServicesRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	var services []db.Service
	var err error
	if req.SearchQuery == "" {
		services, err = server.store.ListServicesByCategory(ctx, req.CategorySlug)
	}
	
	arg := db.ListServicesByNameAndCategoryParams{
		Name:     req.SearchQuery,
		Category: req.CategorySlug,
	}
	
	services, err = server.store.ListServicesByNameAndCategory(ctx, arg)
	
	switch {
	case err != nil:
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	case len(services) == 0:
		ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
	default:
		ctx.JSON(http.StatusOK, services)
	}
}
