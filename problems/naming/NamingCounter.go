package main

import "fmt"

func main() {
	in := []string{"coffee", "donuts", "toffee", "time"}
	fmt.Println("Counting!")
	s := fmt.Sprintf("Result is %d", distinctNames(in))
	fmt.Println(s)
}

func distinctNames(ideas []string) int64 {
	result := int64(0)
	var data = [26]map[string]int{}

	for _, s := range ideas {
		ch := s[0]
		index := ch - 'a'
		rest := s[1:]

		if data[index] == nil {
			data[index] = make(map[string]int)
		}

		if _, ok := data[index][rest]; !ok {
			data[index][rest] = 0
		}
	}

	var pairs [][2]int

	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			if data[i] != nil && data[j] != nil {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}

	for _, pair := range pairs {
		left := data[pair[0]]
		right := data[pair[1]]

		common := 0

		for k1 := range left {
			if _, ok := right[k1]; ok {
				common++
			}
		}

		result += int64(2 * (len(left) - common) * (len(right) - common))
	}
	return result
}
