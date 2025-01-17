package leetcodes

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	t := cleanString(s)
	left := 0
	right := len(t) - 1
	for left < right {
		if t[left] != t[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func cleanString(s string) string {
	var cleaned strings.Builder
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(char)
		}
	}

	return cleaned.String()
}
