package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	n := len(s) - 1
	m := len(t) - 1
	skipCount_s := 0
	skipCount_t := 0
	for n >= 0 || m >= 0 {
		if n >= 0 && s[n] == '#' {
			skipCount_s++
			n--
			continue
		}
		if n >= 0 && skipCount_s > 0 {
			skipCount_s--
			n--
			continue
		}

		if m >= 0 && t[m] == '#' {
			skipCount_t++
			m--
			continue
		}
		if m >= 0 && skipCount_t > 0 {
			skipCount_t--
			m--
			continue
		}

		if n >= 0 && m >= 0 && s[n] != t[m] {
			return false
		}

		n--
		m--
	}
	return n == m
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("b#c", "c"))     // true
	fmt.Println(backspaceCompare("#", "#"))       // true
	fmt.Println(backspaceCompare("#", "a#"))      // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
}
