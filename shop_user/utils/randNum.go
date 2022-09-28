package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandNum(length int) string {
	numarr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	ret := ""
	for i := 0; i < length; i++ {
		ret += strconv.Itoa(numarr[rand.Intn(10)])
	}
	return ret
}
