package api

import (
	"errors"
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern/db/sqlc"
	"github.com/katatrina/SWD392_NET1701_GroupIntern/internal/util"
	"github.com/lib/pq"
)

type createServiceCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	IconURL     string `json:"icon_url" binding:"required"`
	BannerURL   string `json:"banner_url" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// createServiceCategory creates a new service category
//
//	@Router		/service-categories [post]
//	@Summary	Thêm một loại hình dịch vụ
//	@Produce	json
//	@Accept		json
//	@Param		request	body	createServiceCategoryRequest	true	"Service category info"
//	@Description
//	@Tags		service categories
//	@Success	201	{object}	db.ServiceCategory
//	@Failure	400
//	@Failure	500
func (server *Server) createServiceCategory(ctx *gin.Context) {
	var req createServiceCategoryRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	arg := db.CreateServiceCategoryParams{
		Name:        req.Name,
		Slug:        util.Slugify(req.Name),
		IconUrl:     req.IconURL,
		BannerUrl:   req.BannerURL,
		Description: req.Description,
	}
	
	category, err := server.store.CreateServiceCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, category)
}

// listServiceCategories returns a list of service categories
//
//	@Router		/service-categories [get]
//	@Summary	Liệt kê các loại hình dịch vụ
//	@Produce	json
//	@Param		q	query	string	false	"Search query by category name"
//	@Description
//	@Tags		service categories
//	@Success	200	{array}	db.ServiceCategory
//	@Failure	404
//	@Failure	500
func (server *Server) listServiceCategories(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	if searchQuery == "" {
		serviceCategories, err := server.store.ListServiceCategories(ctx)
		switch {
		case err != nil:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		case len(serviceCategories) == 0:
			ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
		default:
			ctx.JSON(http.StatusOK, serviceCategories)
		}
		
		return
	}
	
	serviceCategories, err := server.store.ListServiceCategoriesByName(ctx, searchQuery)
	switch {
	case err != nil:
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	case len(serviceCategories) == 0:
		ctx.JSON(http.StatusNotFound, errorResponse(ErrNoRecordFound))
	default:
		ctx.JSON(http.StatusOK, serviceCategories)
	}
}

// type listServicesByCategoryRequest struct {
// 	CategorySlug string `uri:"slug" binding:"required"`
// }

// listServicesByCategory returns a list of all services of a category
//
//	@Router		/service-categories/{slug}/services [get]
//	@Summary	Liệt kê tất cả dịch vụ của một loại hình
//	@Produce	json
//	@Description
//	@Param		slug	path	string	true	"Category Slug"
//	@Tags		service categories
//	@Success	200	{object}	[]db.Service
//	@Failure	500
// func (server *Server) listServicesByCategory(ctx *gin.Context) {
// 	var req listServicesByCategoryRequest
//
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
//
// 	services, err := server.store.ListServicesByCategory(ctx, req.CategorySlug)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	ctx.JSON(http.StatusOK, services)
// }

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
//	@Tags		service categories
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
//	@Router		/service-categories/{id} [put]
//	@Summary	Cập nhật thông tin của một loại hình dịch vụ
//	@Produce	json
//	@Accept		json
//	@Description
//	@Param		id		path	string							true	"Service Category ID"
//	@Param		request	body	updateServiceCategoryRequest	true	"Update service category info"
//	@Tags		service categories
//	@Success	200
//	@Failure	400
//	@Failure	500
func (server *Server) updateServiceCategory(ctx *gin.Context) {
	var req updateServiceCategoryRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	categoryID, err := server.getLastIDParam(ctx)
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
		IconUrl:     category.IconUrl,
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
//	@Tags		service categories
//	@Success	204
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) deleteServiceCategory(ctx *gin.Context) {
	categoryID, err := server.getLastIDParam(ctx)
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
