package p1004_max_consecutive_ones

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
