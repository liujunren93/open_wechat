package helper

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	baseStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randStr []byte
	buf := []byte(baseStr)
	for i := 0; i < length; i++ {

		randStr = append(randStr, buf[r.Intn(len(baseStr))])
	}
	return string(randStr)
}

func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func RandRange(min, max int) (int, bool) {
	if min > max {
		return 0, false
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intn := r.Intn(max)
	if intn < min {
		return intn + min, true
	}
	return intn, true
}