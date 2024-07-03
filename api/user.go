package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/SWD392_NET1701_GroupIntern/internal/util"
)

var (
	ErrEmailNotFound     = errors.New("email not found")
	ErrPasswordIncorrect = errors.New("password is incorrect")
)

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type userInfo struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

type loginUserResponse struct {
	AccessToken string   `json:"access_token"`
	UserInfo    userInfo `json:"user_info"`
}

// loginUser logs in a user
//
//	@Router		/users/login [post]
//	@Summary	Đăng nhập
//	@Description
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		request	body		loginUserRequest	true	"Login user info"
//	@Success	200		{object}	loginUserResponse
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Failure	500
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	
	// Parse the JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// Get the user by email
	user, err := server.store.GetUserByEmailForLogin(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrEmailNotFound))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	// Check the password
	err = util.CheckPassword(user.HashedPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrPasswordIncorrect))
		return
	}
	
	userID := strconv.FormatInt(user.ID, 10)
	
	// Create a new, unique access token
	accessToken, err := server.tokenMaker.CreateToken(userID, user.Role, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	rsp := loginUserResponse{
		AccessToken: accessToken,
		UserInfo: userInfo{
			ID:          user.ID,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
		},
	}
	
	ctx.JSON(http.StatusOK, rsp)
}
