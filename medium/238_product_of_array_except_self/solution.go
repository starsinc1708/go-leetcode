package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self
func productExceptSelf(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	productExceptSelf()
	fmt.Println(nums)
}
