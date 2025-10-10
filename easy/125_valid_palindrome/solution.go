package main

import (
	"fmt"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
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

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))
}
