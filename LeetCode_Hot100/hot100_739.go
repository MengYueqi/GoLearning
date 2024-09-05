package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	var s []int
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if temperatures[s[j]] < temperatures[i] {
				result[s[j]] = i - s[j]
				s = s[:len(s)-1]
			} else {
				break
			}
		}
		s = append(s, i)
	}
	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
