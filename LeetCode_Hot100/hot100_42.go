package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	sum := 0
	for left < right {
		if leftMax < rightMax {
			left++
			if leftMax >= height[left] {
				sum += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		} else {
			right--
			if rightMax >= height[right] {
				sum += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		}
	}
	return sum
}
