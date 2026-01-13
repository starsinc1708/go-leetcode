package p989_add_to_array_form_to_integer

import (
	"slices"
)

// https://leetcode.com/problems/add-to-array-form-of-integer/
// Time complexity: O(n)
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1
	result := make([]int, 0)
	for k > 0 || i >= 0 {
		if i >= 0 {
			k += num[i]
			i--
		}
		result = append(result, k%10)
		k /= 10
	}
	slices.Reverse(result)
	return result
}
