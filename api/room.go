package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/db/sqlc"
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

type updateRoomRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateRoom updates the name of the room
//
//	@Router		/rooms/{id} [put]
//	@Summary	Cập nhật tên phòng
//	@Description
//	@Params		request body updateRoomRequest true "Update room info"
//	@Param		id	path	int	true	"Room ID"
//	@Tags		rooms
//	@Produce	json
//	@Success	204
//	@Failure	400
//	@Failure	500
func (server *Server) updateRoom(ctx *gin.Context) {
	var req updateRoomRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	roomID, err := server.getLastIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	err = server.store.UpdateRoom(ctx, db.UpdateRoomParams{
		ID:   roomID,
		Name: req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusNoContent, nil)
}
