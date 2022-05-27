package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func SaltFromPassword(pw string) string {
	buff := strings.Builder{}
	for i := 0; i < len(pw); i++ {
		buff.WriteByte(ALPH_UNIVERSE[int(pw[i])%ALPH_UNIVERSE_LEN])
	}
	return buff.String()
}

func HashPassword(pw string) string {
	hash := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(hash[:])
}
