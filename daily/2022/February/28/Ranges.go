package main

import (
	"fmt"
	"strconv"
)

func main() {
	in1 := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(in1))
}

func summaryRanges(nums []int) []string {

	if len(nums) < 1 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	result := []string{}
	start := nums[0]
	curr := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-curr == 1 {
			curr = nums[i]
			continue
		}
		result = append(result, printRange(start, curr))
		curr = nums[i]
		start = nums[i]
	}

	result = append(result, printRange(start, curr))
	return result
}

func printRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}
