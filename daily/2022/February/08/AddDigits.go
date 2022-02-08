package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	x := num % 9

	if x == 0 {
		return 9
	}

	return num % 9
}
