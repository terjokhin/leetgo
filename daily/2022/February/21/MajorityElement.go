package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	length := len(nums)
	result := 0

	for _, n := range nums {
		hash[n]++
		if hash[n] > length/2 {
			result = n
			break
		}

	}
	return result
}
