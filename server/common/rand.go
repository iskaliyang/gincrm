package common

import (
	"math/rand"
	"time"
)

func RandIntSixDigit() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(900000) + 100000
	return randomNumber
}

func RandInt(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min+1) + min
}
