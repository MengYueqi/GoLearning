package LeetCode_Hot100_Round2

// 动态规划有点慢
// func subarraySum(nums []int, k int) int {
// 	sumNum := make([]int, len(nums))
// 	cnt := 0
// 	// 动态规划计算每个可能数组的子和
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j+i < len(nums); j++ {
// 			sumNum[j] = sumNum[j] + nums[j+i]
// 			if sumNum[j] == k {
// 				cnt += 1
// 			}
// 		}
// 	}
// 	return cnt
// }

// 前缀和算法
func subarraySum(nums []int, k int) int {
	// 哈希表保存前缀和
	sumMap := make(map[int]int)
	sumMap[0] = 1
	prevSum, cnt := 0, 0

	for i := 0; i < len(nums); i++ {
		prevSum += nums[i]
		// 如果前缀和中包含 prevSum-k 就有 sumMap[prevSum-k] 个子序列和为 k，进行计数
		if vaule, ok := sumMap[prevSum-k]; ok {
			cnt += vaule
		}
		sumMap[prevSum] += 1
	}

	return cnt

}
