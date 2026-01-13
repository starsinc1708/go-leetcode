package p209_min_size_subarray_sum

import (
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
