package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/../"))
}

func simplifyPath(path string) string {
	var stack []string

	splitted := strings.Split(path, "/")

	for _, el := range splitted {
		if el == "." {
			continue
		}

		if el == ".." {
			if l := len(stack); l > 0 {
				stack = stack[:l-1]
			}
			continue
		}

		if el != "" {
			stack = append(stack, el)
		}
	}

	result := "/"

	for _, el := range stack {
		result += el
		result += "/"
	}

	if l := len(result); l > 1 {
		result = result[:l-1]
	}

	return result
}
