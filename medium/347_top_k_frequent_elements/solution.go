package main

import (
	"fmt"
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

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
}
