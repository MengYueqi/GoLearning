package main

func searchRange(nums []int, target int) []int {
	first, last := 0, len(nums)-1
	idx := len(nums) / 2
	flag := 0
	for first <= last {
		if nums[idx] == target {
			flag = 1
			break
		} else if nums[idx] < target {
			first = idx + 1
			idx = (first + last) / 2
		} else if nums[idx] > target {
			last = idx - 1
			idx = (first + last) / 2
		}
	}
	if flag == 0 {
		return []int{-1, -1}
	}
	first, last = idx-1, idx+1
	for {
		if first >= 0 && nums[first] == target {
			first--
			continue
		} else {
			first++
			break
		}
	}
	for {
		if last < len(nums) && nums[last] == target {
			last++
			continue
		} else {
			last--
			break
		}
	}
	return []int{first, last}
}
