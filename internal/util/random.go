package util

import (
	"math/rand"
)

// RandomIndex returns a random index of a slice
func RandomIndex(limit int) int {
	return rand.Intn(limit) + 1
}
