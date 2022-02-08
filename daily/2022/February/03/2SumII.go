package main

import "fmt"

func main() {

	n1 := []int{1, 2}
	n2 := []int{-2, -1}
	n3 := []int{-1, 2}
	n4 := []int{0, 2}

	fmt.Println(fourSumCount(n1, n2, n3, n4))

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := make(map[int]int)
	result := 0

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			hash[n1+n2]++
		}
	}

	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			result += hash[-(n3 + n4)]
		}
	}
	return result
}
