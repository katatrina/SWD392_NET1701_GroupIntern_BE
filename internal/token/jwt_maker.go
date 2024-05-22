package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateToken(userID string, role string, duration time.Duration) (token string, err error) {
	// Prepare the payload data
	payload, err := NewPayload(userID, role, duration)
	if err != nil {
		return
	}

	// Create an unsigned token with the created payload
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign the token
	token, err = unsignedToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		err = fmt.Errorf("cannot sign the token: %w", err)
		return
	}

	return
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(maker.secretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("token is expired")
		default:
			return nil, fmt.Errorf("token is invalid")
		}
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("unknown payload structure, cannot proceed")
	}

	return payload, nil
}
