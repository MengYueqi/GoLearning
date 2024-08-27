package main

import "fmt"

func maxSubArray(nums []int) int {
	max_value := nums
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if max_value[i-1]+nums[i] > max_value[i] {
			max_value[i] = max_value[i-1] + nums[i]
		}
		if max_value[i] > maxSum {
			maxSum = max_value[i]
		}
	}
	return maxSum
}

func main() {
	fmt.Println("hot100_53")
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
