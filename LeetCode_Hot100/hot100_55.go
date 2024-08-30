package main

func canJump(nums []int) bool {
	farthestPoint := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthestPoint {
			farthestPoint = max(farthestPoint, i+nums[i])
		}
		if farthestPoint >= len(nums)-1 {
			return true
		}
	}
	return false
}
