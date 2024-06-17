package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// listServiceCategories returns a list of all service categories
//
//	@Router		/service-categories [get]
//	@Summary	Liệt kê tất cả danh mục dịch vụ hiện có
//	@Produce	json
//	@Description
//	@Tags		services
//	@Success	200	{object}	[]db.ServiceCategory
//	@Failure	500
func (server *Server) listServiceCategories(ctx *gin.Context) {
	categories, err := server.store.ListServiceCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, categories)
}

type listServicesOfOneCategoryRequest struct {
	CategorySlug string `uri:"slug" binding:"required"`
}

// listServicesOfOneCategory returns a list of all services of a category
//
//	@Router		/service-categories/{slug}/services [get]
//	@Summary	Liệt kê tất cả dịch vụ của một danh mục
//	@Produce	json
//	@Description
//	@Param		slug	path	string	true	"Category Slug"
//	@Tags		services
//	@Success	200	{object}	[]db.Service
//	@Failure	500
func (server *Server) listServicesOfOneCategory(ctx *gin.Context) {
	var req listServicesOfOneCategoryRequest
	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	services, err := server.store.ListServicesOfOneCategory(ctx, req.CategorySlug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, services)
}

type getServiceCategoryBySlugRequest struct {
	CategorySlug string `uri:"slug" binding:"required"`
}

// getServiceCategoryBySlug returns information of a service category
//
//	@Router		/service-categories/{slug} [get]
//	@Summary	Lấy thông tin của một danh mục dịch vụ
//	@Produce	json
//	@Description
//	@Param		slug	path	string	true	"Category Slug"
//	@Tags		services
//	@Success	200	{object}	db.ServiceCategory
//	@Failure	400
//	@Failure	500
func (server *Server) getServiceCategoryBySlug(ctx *gin.Context) {
	var req getServiceCategoryBySlugRequest
	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	category, err := server.store.GetServiceCategoryBySlug(ctx, req.CategorySlug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, category)
}
