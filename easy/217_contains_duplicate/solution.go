package p217_contains_duplicate

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
