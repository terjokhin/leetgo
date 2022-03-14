package main

import "fmt"

func main() {

	fmt.Println(quickSort([]int{1, 6, 3, 4, 78, 9, 9, 10}))

}

func quickSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	var left, right []int

	for _, n := range nums[1:] {
		if n < pivot {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}

	sortedLeft :=  quickSort(left)
	sortedRight := quickSort(right)

	result := append(sortedLeft, pivot)
	result = append(result, sortedRight...)

	return result
}
