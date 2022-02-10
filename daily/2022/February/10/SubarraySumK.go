package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {

	hashMap := make(map[int]int)
	results := 0
	sum := 0

	hashMap[0] = 1

	for _, n := range nums {
		sum += n
		if v, exists := hashMap[sum-k]; exists {
			results += v
		}
		hashMap[sum]++
	}
	return results
}
