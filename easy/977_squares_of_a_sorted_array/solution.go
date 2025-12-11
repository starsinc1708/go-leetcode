package main

import "fmt"

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

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
