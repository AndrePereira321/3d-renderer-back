package utils

import "crypto/sha256"

func SaltFromPassword(pw string) string {
	salt := ""
	for i := 0; i < len(pw); i++ {
		salt += string(ALPH_UNIVERSE[int(pw[i])%ALPH_UNIVERSE_LEN])
	}
	return salt
}

func HashPassword(pw string) string {
	hash := sha256.Sum256([]byte(pw))
	return string(hash[:])
}
