package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// listAllServiceCategories returns a list of all service categories
//
//	@Router		/service-categories [get]
//	@Summary	list all service categories
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
