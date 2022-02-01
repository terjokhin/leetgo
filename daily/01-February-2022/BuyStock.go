package main

//You are given an array prices where prices[i] is the price of a given stock on the ith day.
//You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

import "fmt"

func main() {

	example1 := []int{7, 1, 5, 3, 6, 4}
	example2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(example1))
	fmt.Println(maxProfit(example2))
}

func maxProfit(prices []int) int {
	currMin := -1
	currV := 0

	for _, n := range prices {
		if currMin == -1 {
			currMin = n
			continue
		}

		if n > currMin && currV < n-currMin {
			currV = n - currMin
		} else if n < currMin {
			currMin = n
		}

	}
	return currV
}
