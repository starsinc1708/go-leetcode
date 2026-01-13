package p242_valid_anagram

// https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countM := map[rune]int{}

	for _, r := range s {
		countM[r]++
	}

	for _, r := range t {
		if countM[r] == 0 {
			return false
		}
		countM[r]--
	}

	return true
}
