package utils

import (
	"math/rand"
	"strings"
)

func RandomNumber(min, max int) int {
	if min < max {
		return min + rand.Intn(max-min)
	}
	return min + rand.Intn(min-max)
}

func RandomString(len int) string {
	buff := strings.Builder{}
	for i := 0; i < len; i++ {
		buff.WriteByte(ALPH_UNIVERSE[rand.Intn(ALPH_UNIVERSE_LEN)])
	}
	return buff.String()
}
