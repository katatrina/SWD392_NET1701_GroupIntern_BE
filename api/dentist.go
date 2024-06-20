package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type listDentistsResponse struct {
}

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
