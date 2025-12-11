package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/merge-sorted-array
// Solution Complexity:
// Time: O((n+m)*log(n+m))
// Space: O(1)
func mergeWithSort(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func mergeWithCopy(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

// O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	result := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[result] = nums1[p1]
			p1--
		} else {
			nums1[result] = nums2[p2]
			p2--
		}
		result--
	}

	for p2 >= 0 {
		nums1[result] = nums2[p2]
		p2--
		result--
	}

	fmt.Println(nums1)
}

func main() {
	mergeWithSort([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	mergeWithCopy([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
