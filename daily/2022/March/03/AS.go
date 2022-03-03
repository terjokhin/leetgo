package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1}))
}

func numberOfArithmeticSlices(nums []int) int {

	if len(nums) < 3 {
		return 0
	}

	result := 0
	acc := 0

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			acc++
			result += acc
		} else {
			acc = 0
		}
	}
	return result

}
