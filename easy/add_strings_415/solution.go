package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/add-strings
// Time complexity: O(max(m,n))
// Space complexity: O(max(m,n))
func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	idx1, idx2 := len(num1)-1, len(num2)-1
	reminder := 0
	var result []byte

	for idx1 >= 0 || idx2 >= 0 {
		digit1, digit2 := 0, 0

		if idx1 >= 0 {
			digit1 = int(num1[idx1] - '0')
		}
		if idx2 >= 0 {
			digit2 = int(num2[idx2] - '0')
		}

		reminder += digit1 + digit2
		result = append(result, byte(reminder%10)+'0')

		reminder /= 10
		idx1--
		idx2--
	}

	if reminder != 0 {
		result = append(result, byte(reminder%10)+'0')
	}

	slices.Reverse(result)

	return string(result)
}

func main() {
	fmt.Println(addStrings("456", "77"))
}
