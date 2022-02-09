package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 1, 1, 1, 1}, 0))
}

func findPairs(nums []int, k int) int {
	results := 0
	hashMap := make(map[int]int)

	for _, n := range nums {
		hashMap[n]++
	}

	if k >= 0 {

		for n, v := range hashMap {
			if k == 0 && v > 1 {
				results++
				continue
			}

			if k > 0 {
				_, exists := hashMap[n+k]
				if exists {
					results++
				}
			}

		}
	}

	return results
}
