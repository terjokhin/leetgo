package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6))
}

func binarySearch(input []int, target int) int {
	from, to := 0, len(input)-1
	mid := (to - from) / 2

	for input[mid] != target {
		fmt.Println(".")
		if input[mid] > target {
			to = mid
		} else {
			from = mid

		}
		mid = (to + from) / 2
	}
	return mid
}
