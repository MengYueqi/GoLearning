package main

func jump(nums []int) int {
	cnt, last, lastIdx := 0, 0, 0
	for lastIdx < len(nums)-1 {
		for i := 0; i <= lastIdx; i++ {
			last = max(last, i+nums[i])
		}
		lastIdx = last
		cnt++
	}
	return cnt
}
