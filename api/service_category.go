package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/util"
	"github.com/lib/pq"
)

// listServiceCategories returns a list of all service categories
//
//	@Router		/service-categories [get]
//	@Summary	Liệt kê tất cả loại hình dịch vụ hiện có
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
//	@Summary	Liệt kê tất cả dịch vụ của một loại hình dịch vụ
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
//	@Summary	Lấy thông tin của một loại hình dịch vụ
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

type updateServiceCategoryRequest struct {
	Name        *string `json:"name"`
	IconURL     *string `json:"icon_url"`
	BannerURL   *string `json:"banner_url"`
	Description *string `json:"description"`
}

// updateServiceCategory updates information of a service category
//
//	@Router		/service-categories/{id} [patch]
//	@Summary	Cập nhật thông tin của một loại hình dịch vụ
//	@Produce	json
//	@Accept		json
//	@Description
//	@Param		id		path	string							true	"Service Category ID"
//	@Param		request	body	updateServiceCategoryRequest	true	"Update service category info"
//	@Tags		services
//	@Success	200
//	@Failure	400
//	@Failure	500
func (server *Server) updateServiceCategory(ctx *gin.Context) {
	var req updateServiceCategoryRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	categoryID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	category, err := server.store.GetServiceCategoryByID(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if req.Name != nil {
		category.Name = *req.Name
		category.Slug = util.Slugify(*req.Name)
	}
	if req.IconURL != nil {
		category.IconUrl = *req.IconURL
	}
	if req.BannerURL != nil {
		category.BannerUrl = *req.BannerURL
	}
	if req.Description != nil {
		category.Description = *req.Description
	}
	
	arg := db.UpdateServiceCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		BannerUrl:   category.BannerUrl,
		Description: category.Description,
	}
	err = server.store.UpdateServiceCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, nil)
}

// deleteServiceCategory deletes a service category
//
//	@Router		/service-categories/{id} [delete]
//	@Summary	Xóa một loại hình dịch vụ
//	@Description
//	@Param		id	path	int	true	"Service Category ID"
//	@Tags		services
//	@Success	204
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) deleteServiceCategory(ctx *gin.Context) {
	categoryID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	err = server.store.DeleteServiceCategory(ctx, categoryID)
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
