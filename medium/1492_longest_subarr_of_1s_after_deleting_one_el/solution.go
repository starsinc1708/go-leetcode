package p1492_longest_subarr_of_1s_after_deleting_one_el

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
