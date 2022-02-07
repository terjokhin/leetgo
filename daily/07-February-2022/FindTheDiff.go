package main

import "fmt"

func main() {
	original := "abcd"
	poddelka := "abcde"
	fmt.Println(string(findTheDifference(original, poddelka)))
}

func findTheDifference(s string, t string) byte {

	indexes := [26]int{}

	for _, n := range s {
		indexes[n-'a']--
	}

	for _, n := range t {
		indexes[n-'a']++
	}

	var result byte

	for i, e := range indexes {
		if e > 0 {
			result = byte('a' + i)
			break
		}
	}

	return result
}
