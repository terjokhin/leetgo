package main

import "fmt"

func main() {

	fmt.Println(findAnagrams("abab", "ab"))

}

func findAnagrams(s string, p string) []int {
	var pArr, sArr [26]int
	pLen := len(p)
	var results []int

	if len(s) < len(p) {
		return results
	}

	for _, pCh := range p {
		pArr[pCh-'a']++
	}

	for i := range p {
		sArr[s[i]-'a']++
	}

	if pArr == sArr {
		results = append(results, 0)
	}

	for i := pLen; i < len(s); i++ {
		sArr[s[i-pLen]-'a']--
		sArr[s[i]-'a']++
		if pArr == sArr {
			results = append(results, i-pLen+1)
		}
	}

	return results
}
