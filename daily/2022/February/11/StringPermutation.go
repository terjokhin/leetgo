package main

import "fmt"

func main() {

	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))

}

func checkInclusion(s1 string, s2 string) bool {

	exLen := len(s1)
	crLen := len(s2)

	if exLen > crLen {
		return false
	}

	var example, current [26]int

	for _, ch := range s1 {
		example[ch-'a']++
	}

	for i := 0; i < exLen; i++ {
		current[s2[i]-'a']++
	}

	if example == current {
		return true
	}

	for i := exLen; i < crLen; i++ {
		current[s2[i-exLen]-'a']--
		current[s2[i]-'a']++
		if example == current {
			return true
		}
	}
	return false
}
