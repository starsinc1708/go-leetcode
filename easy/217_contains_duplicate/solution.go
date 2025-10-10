package main

import "fmt"

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	m := map[int]struct{}{}
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5}))
	fmt.Println(containsDuplicate([]int{1, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
