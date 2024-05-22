package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	jwt.RegisteredClaims
	Role string `json:"role,omitempty"`
}

func NewPayload(userID string, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot create token id: %w", err)
	}

	payload := &Payload{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chauvinhphuoc",
			Subject:   userID,
			Audience:  []string{"client"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        tokenID.String(),
		},
		Role: role,
	}

	return payload, nil
}
