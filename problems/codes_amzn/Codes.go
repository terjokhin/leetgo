package main

import (
	"fmt"
	"strings"
)

func main() {

	codes := []string{"apple banana", "orange anything banana"}
	cart := []string{"apple", "banana", "apple", "orange", "potato", "banana"}

	fmt.Println(run(codes, cart))
}

func run(codes []string, cart []string) int {
	if len(codes) == 0 {
		return 0
	}

	var res int

	var codesArray [][]string

	for _, n := range codes {
		codesSplited := strings.Split(n, " ")
		codesArray = append(codesArray, codesSplited)
		res += len(codesSplited)
	}

	matches := 0
	smatches := 0

	clp := 0
	cip := 0
	currC := codesArray[clp]
	clp++

	for _, p := range cart {

		if cip >= len(currC) {
			matches += smatches
			smatches = 0
			cip = 0
			currC = codesArray[clp]
			clp++
		}

		if match(currC[cip], p) {
			smatches++
			cip++
		} else {
			cip = 0
			smatches = 0
		}
	}

	if clp != len(codes) {
		return 0
	}

	matches += smatches

	if matches >= res {
		return 1
	}

	return 0
}

func match(code, value string) bool {
	if code == "anything" {
		return true
	}

	return code == value
}
