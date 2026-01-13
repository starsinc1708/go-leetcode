package p844_backspace_string_compare

func backspaceCompare(s string, t string) bool {
	n := len(s) - 1
	m := len(t) - 1
	skipCount_s := 0
	skipCount_t := 0
	for n >= 0 || m >= 0 {
		if n >= 0 && s[n] == '#' {
			skipCount_s++
			n--
			continue
		}
		if n >= 0 && skipCount_s > 0 {
			skipCount_s--
			n--
			continue
		}

		if m >= 0 && t[m] == '#' {
			skipCount_t++
			m--
			continue
		}
		if m >= 0 && skipCount_t > 0 {
			skipCount_t--
			m--
			continue
		}

		if n >= 0 && m >= 0 && s[n] != t[m] {
			return false
		}

		n--
		m--
	}
	return n == m
}
