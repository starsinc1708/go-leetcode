package main

import "fmt"

func isMatching(open rune, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '{' && close == '}') ||
		(open == '[' && close == ']')
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !isMatching(prev, v) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
}
