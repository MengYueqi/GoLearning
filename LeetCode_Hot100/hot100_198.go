package main

func rob(nums []int) int {
	sum := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else {
		sum[0], sum[1] = nums[0], max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			sum[i] = max(nums[i]+sum[i-2], sum[i-1])
		}
	}
	return sum[len(nums)-1]
}
