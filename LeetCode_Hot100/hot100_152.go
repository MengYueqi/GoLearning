package main

func maxProduct(nums []int) int {
	maxF, minF, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax, tempMin := maxF, minF
		maxF = max(nums[i], nums[i]*tempMax, nums[i]*tempMin)
		minF = min(nums[i], nums[i]*tempMax, nums[i]*tempMin)
		result = max(result, maxF)
	}
	return result
}
