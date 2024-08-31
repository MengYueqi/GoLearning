package main

func orderSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	first, last := 0, len(nums)-1
	idx := len(nums) / 2
	for first <= last {
		if nums[idx] == target {
			return idx
		} else if nums[idx] < target {
			first = idx + 1
			idx = (first + last) / 2
		} else {
			last = idx - 1
			idx = (first + last) / 2
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	} else if len(nums) == 0 {
		return -1
	}
	isOrder1 := false
	idx := len(nums) / 2
	if nums[idx] > nums[0] {
		isOrder1 = true
	} else {
		isOrder1 = false
	}
	if isOrder1 {
		if target == nums[idx] {
			return idx
		} else if target < nums[idx] && target >= nums[0] {
			return orderSearch(nums[:idx], target)
		} else {
			if search(nums[idx+1:], target) == -1 {
				return -1
			} else {
				return idx + search(nums[idx+1:], target) + 1
			}
		}
	} else {
		if target == nums[idx] {
			return idx
		} else if target > nums[idx] && target <= nums[len(nums)-1] {
			if orderSearch(nums[idx+1:], target) == -1 {
				return -1
			} else {
				return idx + orderSearch(nums[idx+1:], target) + 1
			}
		} else {
			return search(nums[0:idx], target)
		}
	}
}
