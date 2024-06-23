package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type createRoomRequest struct {
	Name string `json:"name" binding:"required"`
}

// createRoom creates a new room
//
//	@Router		/rooms [post]
//	@Summary	Tạo một phòng mới
//	@Description
//	@Tags		rooms
//	@Accept		json
//	@Produce	json
//	@Param		request	body		createRoomRequest	true	"Create room info"
//	@Success	201		{object}	db.Room
//	@Failure	400
//	@Failure	500
func (server *Server) createRoom(ctx *gin.Context) {
	var req createRoomRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	room, err := server.store.CreateRoom(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusCreated, room)
}

// listRooms returns a list of rooms
//
//	@Router		/rooms [get]
//	@Summary	Lấy danh sách tất cả phòng
//	@Description
//	@Tags		rooms
//	@Produce	json
//	@Success	200	{object}	[]db.Room
//	@Failure	500
func (server *Server) listRooms(ctx *gin.Context) {
	rooms, err := server.store.ListRooms(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, rooms)
}
