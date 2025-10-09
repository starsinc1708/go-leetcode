package main

import (
	"sort"
)

// https://leetcode.com/problems/merge-sorted-array
// Solution Complexity:
// Time: O(n log n)
// Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
