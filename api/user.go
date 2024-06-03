package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/util"
	"github.com/lib/pq"
)

var (
	ErrEmailNotFound     = errors.New("email not found")
	ErrPasswordIncorrect = errors.New("password is incorrect")
)

type createCustomerRequest struct {
	Password    string `json:"password" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// createCustomer creates a new customer
//
//	@Router		/users [post]
//	@Summary	create a new customer
//	@Description
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		customer	body	createCustomerRequest	true	"Create customer"
//	@Success	201
//	@Failure	400
//	@Failure	403
//	@Failure	500
func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomerRequest

	// Parse the JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Generate hashed password
	hashedPassword, err := util.GenerateHashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateCustomerParams{
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
	}

	// Create a new customer
	_, err = server.store.CreateCustomer(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch {
			case pqErr.Code.Name() == "unique_violation":
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

	ctx.JSON(http.StatusCreated, nil)
}

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
//	@Summary	log in a user
//	@Description
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user	body		loginUserRequest	true	"Login user"
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
	user, err := server.store.GetUserByEmail(ctx, req.Email)
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
	accessToken, err := server.tokenMaker.CreateToken(userID, user.Role, time.Minute*60)
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
