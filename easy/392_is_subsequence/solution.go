package main

import "fmt"

func isSubsequence(s string, t string) bool {
	p1, p2 := 0, 0
	for ; p1 < len(t) && p2 < len(s); p1++ {
		if t[p1] == s[p2] {
			p2++
		}
	}
	return p2 == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "adsfxdexc"))
}
