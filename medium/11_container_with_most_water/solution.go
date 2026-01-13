package p11_container_with_most_water

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		minHeight := min(height[left], height[right])

		maxArea = max(maxArea, minHeight*width)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
