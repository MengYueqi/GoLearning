package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxA := 0
	low := 0
	high := len(height) - 1
	for low < high {
		// 更新最大值
		if (high-low)*min(height[low], height[high]) > maxA {
			maxA = (high - low) * min(height[low], height[high])
		}
		if height[low] <= height[high] {
			low += 1
		} else if height[low] > height[high] {
			high -= 1
		}
	}
	return maxA
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
