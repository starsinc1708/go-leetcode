package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	begin := 0
	windowState := 0
	result := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			windowState++
		}

		for windowState > k {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}

		result = max(result, end-begin+1)
	}

	return result
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
