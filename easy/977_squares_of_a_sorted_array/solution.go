package p977_squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	left, right := 0, len(nums)-1

	i := len(nums) - 1

	for i >= 0 {
		squareL := nums[left] * nums[left]
		squareR := nums[right] * nums[right]
		if squareL < squareR {
			result[i] = squareR
			right--
		} else {
			result[i] = squareL
			left++
		}
		i--
	}
	return result
}
