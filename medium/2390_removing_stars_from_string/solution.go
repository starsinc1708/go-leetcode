package main

import "fmt"

func removeStars(s string) string {
	var stack []uint8

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func main() {
	fmt.Println(removeStars("***"))         // ""
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("eerase*****")) // "e"
	fmt.Println(removeStars("erase*****1")) // "1"

}
