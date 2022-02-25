package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	ans := 0
	for _, n := range columnTitle {
		ans = 26*ans + (int(n) - 'A' + 1)
	}
	return ans
}
