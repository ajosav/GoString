package str

import (
	"math/rand"
	"strings"
)

func Squish(string string) string {
	return strings.Join(strings.Fields(string), " ")
}

func Mask(s string, maskChar rune, revealCount int) string {
	if revealCount <= 0 {
		return strings.Repeat(string(maskChar), len(s))
	}

	if revealCount >= len(s) {
		return s
	}

	revealedPart := s[:revealCount]
	maskedLength := len(s) - revealCount
	return revealedPart + strings.Repeat(string(maskChar), maskedLength)
}

func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
