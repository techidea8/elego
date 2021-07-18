package utils

import (
	"math/rand"
	"time"

	"github.com/go-basic/uuid"
)

const randomseedstr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandomString(l int) string {
	var container string
	l1 := len(randomseedstr)
	for i := 0; i < l; i++ {
		randomInt := rand.Intn(l1)
		container += string(randomseedstr[randomInt])
	}
	return container
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func UuidString() string {
	return uuid.New()
}
