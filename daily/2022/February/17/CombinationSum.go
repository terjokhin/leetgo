package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	splitBy := -1
	for i, n := range candidates {
		if n > target {
			splitBy = i
			break
		}
	}
	if splitBy > 0 {
		candidates = candidates[:splitBy]
	}
	var results [][]int
	run(candidates, target, []int{}, &results)
	return results
}

func run(candidates []int, target int, curr []int, results *[][]int) {
	for i, next := range candidates {
		if target-next <= 0 {
			if target-next == 0 {
				// must make a copy of current for result
				match := make([]int, len(curr)+1)
				copy(match, curr)
				match[len(curr)] = next
				*results = append(*results, match)
			}
			break
		}
		run(candidates[i:], target-next, append(curr, next), results)
	}
}
