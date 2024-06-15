package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// listAllServiceCategories returns a list of all service categories
//
//	@Router		/service-categories [get]
//	@Summary	liệt kê tất cả danh mục dịch vụ hiện có
//	@Produce	json
//	@Description
//	@Tags		services
//	@Success	200 {object} []db.ServiceCategory
//	@Failure	500
func (server *Server) listAllServiceCategories(ctx *gin.Context) {
	categories, err := server.store.ListAllServiceCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, categories)
}

type listAllServicesOfCategoryRequest struct {
	CategoryID int64 `uri:"id" binding:"required"`
}

// listAllServicesOfACategory returns a list of all services of a category
//
//	@Router		/service-categories/{id}/services [get]
//	@Summary	liệt kê tất cả dịch vụ của một danh mục
//	@Produce	json
//	@Description
//	@Param		id	path	int	true	"Category ID"
//	@Tags		services
//	@Success	200 {object} []db.Service
//	@Failure	500
func (server *Server) listAllServicesOfACategory(ctx *gin.Context) {
	var req listAllServicesOfCategoryRequest
	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	services, err := server.store.ListAllServicesOfACategory(ctx, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, services)
}
