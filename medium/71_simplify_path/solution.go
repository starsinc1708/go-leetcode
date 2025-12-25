package main

import (
	"fmt"
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

func main() {
	fmt.Println(simplifyPath("/home/"))                           // "/home"
	fmt.Println(simplifyPath("/home//foo/"))                      // "/home/foo"
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures")) // "/home/user/Pictures"
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))            // "/.../b/d"
}
