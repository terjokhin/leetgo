package main

import "fmt"

func main() {
	// s1 := "abc"
	// t1 := "ahbgdc"
	// s2 := "axc"
	// t2 := "ahbgdc"
	s3 := "aaaaaa"
    t3 := "bbaaaa"
	//fmt.Println(isSubsequence(s1, t1))
	//fmt.Println(isSubsequence(s2, t2))
	fmt.Println(isSubsequence(s3, t3))
}

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	found := false
	position := 0

	for _, ch := range s {
		found = false

		for i := position; i < len(t); i++ {
			if rune(t[i]) == ch {
				found = true
				position = i + 1
				break
			}
		}
		if !found {
			return found
		}
	}
	return found
}
