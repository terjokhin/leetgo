package main

import "fmt"

func main() {

	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	result := 0

	for _, n := range nums {
		r := hash[n]
		if r != 0 {
			delete(hash, n)
		} else {
			hash[n]++
		}
	}

	for k := range hash {
		result = k
	}

	return result
}
