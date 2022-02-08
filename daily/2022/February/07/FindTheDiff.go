package main

import "fmt"

func main() {
	original := "abcd"
	poddelka := "abcde"
	fmt.Println(string(findTheDifference(original, poddelka)))
}

func findTheDifference(s string, t string) byte {
	var result int32

	for _, n := range s {
		result ^= n
	}

	for _, n := range t {
		result ^= n
	}

	return byte(result)
}
