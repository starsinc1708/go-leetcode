package p71_simplify_path

import (
	"strings"
)

// if .. remove from stack

func simplifyPath(path string) string {
	var stack []string

	for _, dir := range strings.Split(path, "/") {
		if dir == "" || dir == "." {
			continue
		}
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, dir)
		}
	}

	return "/" + strings.Join(stack, "/")
}
