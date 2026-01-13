package p125_valid_palindrome

import (
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/
func IsPalindrome(s string) bool {
	var pStr []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			pStr = append(pStr, unicode.ToLower(r))
		}
	}

	l := 0
	r := len(pStr) - 1

	for l <= r {
		if pStr[l] != pStr[r] {
			return false
		}
		l++
		r--
	}

	return true
}
