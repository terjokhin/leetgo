package main

import "fmt"

func main() {

	example1 := []int{0, 1}

	fmt.Println(findMaxLength(example1))
}

func findMaxLength(nums []int) int {

	var count int
	var max int
	countMap := make(map[int]int)
	countMap[0] = -1

	for i, n := range nums {

		if n == 1 {
			count++
		} else {
			count--
		}

		if v, exists := countMap[count]; exists {
			if i-v > max {
				max = i - v
			}
		} else {
			countMap[count] = i
		}

	}

	return max
}
