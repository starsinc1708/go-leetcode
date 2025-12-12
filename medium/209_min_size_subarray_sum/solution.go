package main

import (
	"fmt"
	"math"
)

// O(2n) worst
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	begin := 0
	windowState := 0
	result := math.MaxInt64

	for end := 0; end < len(nums); end++ {
		windowState += nums[end]

		for windowState >= target {
			windowSize := end - begin + 1
			result = min(result, windowSize)
			windowState -= nums[begin]
			begin++
		}
	}

	if result == math.MaxInt64 {
		return 0
	}

	return result
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))  // 2
	fmt.Println(minSubArrayLen(11, []int{2, 3, 1, 2, 4, 3})) // 5
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1}))          // 0
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))           // 1

}
