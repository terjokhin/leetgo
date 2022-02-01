package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	currentStart := 0
	currentStop := 0

	for i := range s {

		//for odd
		i1 := i
		j1 := i
		for i1 >= 0 && j1 < len(s) && s[i1] == s[j1] {
			if currentStop-currentStart < j1-i1 {
				currentStart = i1
				currentStop = j1
			}
			i1--
			j1++
		}

		//for even
		i1 = i
		j1 = i + 1
		for i1 >= 0 && j1 < len(s) && s[i1] == s[j1] {
			if currentStop-currentStart < j1-i1 {
				currentStart = i1
				currentStop = j1
			}
			i1--
			j1++
		}
	}
	return s[currentStart : currentStop+1]
}
