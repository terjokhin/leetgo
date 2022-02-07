package main

import "fmt"

func main() {

	i := 121
	j := -121

	fmt.Println(isPalindrome(i))
	fmt.Println(isPalindrome(j))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	reversed := 0

	for y := x; y > 0; y /= 10 {
		reversed = reversed*10 + y%10
	}

	return x == reversed
}
