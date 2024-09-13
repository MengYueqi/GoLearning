package main

func swap(nums []int, idx int) {
	for i := idx; i <= (idx+len(nums)-1)/2; i++ {
		nums[i], nums[len(nums)-1+idx-i] = nums[len(nums)-1+idx-i], nums[i]
	}
}

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j >= 0; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					swap(nums, i+1)
					return
				}
			}
		}
	}
	swap(nums, 0)
	return
}
