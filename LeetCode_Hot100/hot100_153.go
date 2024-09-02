package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return min(nums[0], nums[1])
	}
	idx := len(nums) / 2
	if nums[idx] < nums[0] {
		return findMin(nums[0 : idx+1])
	} else if nums[idx] > nums[len(nums)-1] {
		fmt.Print("Here")
		return findMin(nums[idx+1:])
	} else {
		return nums[0]
	}
}

func main() {
	fmt.Println(findMin([]int{2, 1}))
}
