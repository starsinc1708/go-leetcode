package p26_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	left := 0
	right := 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}
