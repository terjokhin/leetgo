package main

import "fmt"

func main() {

	fmt.Println(getSmallestString(3, 27))
	fmt.Println(getSmallestString(5, 73))

}

func getSmallestString(n int, k int) string {

	result := make([]byte, n)
	nextValue := 26

	for n > 0 && k != n {
		n--
		nextValue = min(26, k-n)
		result[n] = byte('a' - 1 + nextValue)
		k = k - nextValue
	}

	for n > 0 {
		n--
		result[n] = 'a'
	}

	return string(result[:])
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
