package main

import "fmt"

// https://leetcode.com/problems/two-sum/
// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	m := make(map[int]int)

	for i, num := range nums {
		reminder := target - num
		index, found := m[reminder]
		if found {
			return []int{index, i}
		}
		m[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 1, 5}, 6))
}
