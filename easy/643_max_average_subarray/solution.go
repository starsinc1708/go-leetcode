package p643_max_average_subarray

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	windowState := float64(0)
	result := math.Inf(-1)
	for end := range nums {
		windowState += float64(nums[end])
		if end-begin+1 == k {
			result = math.Max(result, windowState)
			windowState -= float64(nums[begin])
			begin++
		}
	}

	return result / float64(k)
}
