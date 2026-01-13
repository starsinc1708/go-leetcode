package p1047_remove_all_adjacent_duplicates_in_string

func removeDuplicates(s string) string {
	stack := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == rune(s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, rune(s[i]))
		}
	}
	return string(stack)
}
