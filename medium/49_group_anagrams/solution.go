package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		sortedS := sortString(s)
		groups[sortedS] = append(groups[sortedS], s)
	}

	var result [][]string

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "nat", "bat"}))
}
