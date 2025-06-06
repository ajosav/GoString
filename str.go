package str

import "strings"

func Squish(string string) string {
	return strings.Join(strings.Fields(string), "")
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
