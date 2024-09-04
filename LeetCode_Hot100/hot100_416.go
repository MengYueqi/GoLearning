package main

import "fmt"

func canPartition(nums []int) bool {
	sumOfNums := 0
	for i := 0; i < len(nums); i++ {
		sumOfNums += nums[i]
	}
	if sumOfNums%2 == 1 {
		return false
	}

	// 初始化 dp 记录矩阵
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, (sumOfNums/2)+1)
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], nums[i-1]+dp[i-1][j-nums[i-1]])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(nums)][sumOfNums/2] == sumOfNums/2
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
