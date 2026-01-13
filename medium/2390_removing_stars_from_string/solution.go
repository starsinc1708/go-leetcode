package p2390_removing_stars_from_string

func removeStars(s string) string {
	var stack []uint8

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
