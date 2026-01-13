package p238_product_of_array_except_self

// https://leetcode.com/problems/product-of-array-except-self
func productExceptSelf(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	return nums
}
