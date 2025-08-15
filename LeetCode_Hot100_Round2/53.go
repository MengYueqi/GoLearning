package LeetCode_Hot100_Round2

// 动态规划算法时间复杂度较高
// func maxSubArray(nums []int) int {
// 	sum := make([]int, len(nums))
// 	for i := 0; i < len(sum); i++ {
// 		sum[i] = 0
// 	}
// 	maxValue := nums[0]
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums); j++ {
// 			sum[j-i] = sum[j-i] + nums[j]
// 			maxValue = max(maxValue, sum[j-i])
// 		}
// 	}
// 	return maxValue
// }

// 贪心算法
func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum = curSum + nums[i]
		}
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
