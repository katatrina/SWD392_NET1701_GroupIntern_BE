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
	ErrIncorrectEmailOrPassword = errors.New("email or password is incorrect")
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
//	@Param		customer	body		createCustomerRequest	true	"Create customer"
//	@Success	201			{object}	nil
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

type userResponse struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	UserInfo    userResponse `json:"user_info"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// Parse the JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the customer by email
	customer, err := server.store.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrIncorrectEmailOrPassword))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check the password
	err = util.CheckPassword(customer.HashedPassword, req.Password)
	if err != nil {
		err = errors.New("incorrect email or password")
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrIncorrectEmailOrPassword))
		return
	}

	userID := strconv.FormatInt(customer.ID, 10)

	// Create a new, unique access token
	accessToken, err := server.tokenMaker.CreateToken(userID, customer.Role, time.Minute*15)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		UserInfo: userResponse{
			ID:          customer.ID,
			FullName:    customer.FullName,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
			Role:        customer.Role,
			CreatedAt:   customer.CreatedAt,
		},
	}
	ctx.JSON(http.StatusOK, rsp)
}
