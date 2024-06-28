package api

import (
	"errors"
	"net/http"
	"strings"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/SWD392_NET1701_GroupIntern/internal/token"
)

var (
	ErrAuthorizationHeaderNotProvided   = errors.New("authorization header is not provided")
	ErrInvalidAuthorizationHeaderFormat = errors.New("invalid authorization header format")
	ErrAuthorizationTypeNotSupported    = errors.New("authorization type is not supported")
)

var (
	authorizationPayloadKey = "authorization_payload"
	authorizationHeaderType = "bearer"
)

// authMiddleware is a middleware to check if the request is authorized.
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrAuthorizationHeaderNotProvided))
			return
		}
		
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrInvalidAuthorizationHeaderFormat))
			return
		}
		
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationHeaderType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrAuthorizationTypeNotSupported))
			return
		}
		
		accessToken := fields[1]
		claims, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		
		ctx.Set(authorizationPayloadKey, claims)
		ctx.Next()
	}
}
