package main

import "fmt"

func longestSubarray(nums []int) int {
	k := 1

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

	return result - 1
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                // 3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) // 5
}
