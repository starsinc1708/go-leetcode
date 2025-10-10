package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countM := map[rune]int{}

	for _, r := range s {
		countM[r]++
	}

	for _, r := range t {
		if countM[r] == 0 {
			return false
		}
		countM[r]--
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("rat", "art"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
