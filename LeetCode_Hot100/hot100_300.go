package main

func lengthOfLIS(nums []int) int {
	cnt := make([]int, len(nums))
	for i := 0; i < len(cnt); i++ {
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				cnt[i] = max(cnt[i], cnt[j]+1)
			}
		}
	}
	maxNum := cnt[0]
	for i := 0; i < len(cnt); i++ {
		maxNum = max(maxNum, cnt[i])
	}
	return maxNum
}
