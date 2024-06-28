package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// listSpecialties returns a list of specialties
//
//	@Router		/specialties [get]
//	@Summary	Liệt kê tất cả chuyên khoa
//	@Produce	json
//	@Description
//	@Tags		specialties
//	@Success	200 {array} db.Specialty
//	@Failure	400 
//	@Failure	500
func (server *Server) listSpecialties(ctx *gin.Context) {
	// Get all specialties
	specialties, err := server.store.ListSpecialties(ctx)
	// Check if any error is returned
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Send an array of specialties as the JSON response to the client
	ctx.JSON(http.StatusOK, specialties)
}
