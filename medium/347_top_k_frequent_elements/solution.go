package p347_top_k_frequent_elements

import (
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-elements
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{nums[0]}
	}

	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	type kv struct {
		Key   int
		Value int
	}

	var pairs []kv
	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	topKeys := make([]int, k)

	for i := 0; i < k; i++ {
		topKeys[i] = pairs[i].Key
	}

	return topKeys
}
