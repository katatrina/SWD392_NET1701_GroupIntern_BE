package token

import (
	"time"
)

type Maker interface {
	CreateToken(userID string, role string, duration time.Duration) (token string, err error)

	VerifyToken(token string) (payload *Payload, err error)
}
