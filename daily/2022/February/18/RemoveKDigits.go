package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("112", 1))
}

func removeKdigits(num string, k int) string {
	length := len(num)
	stack := make([]string, 0)

	if k >= length {
		return "0"
	}
	i := 0
	for i < length && k > 0 {
		if len(stack) > 0 && stack[len(stack)-1] > string(num[i]) {
			stack = stack[:len(stack)-1]
			k--
		} else {
			stack = append(stack, string(num[i]))
			i++
		}
	}
	for i < length {
		stack = append(stack, string(num[i]))
		i++
	}

	if k != 0 {
		stack = stack[:len(stack)-k]
	}

	for len(stack) > 1 && stack[0] == "0" {
		stack = stack[1:]
	}

	return strings.Join(stack, "")
}
