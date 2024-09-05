package main

func majorityElement(nums []int) int {
	cntMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cntMap[nums[i]] += 1
	}
	result, maxCnt := 0, 0
	for key, value := range cntMap {
		if value >= len(nums)/2 && value > maxCnt {
			maxCnt = value
			result = key
		}
	}
	return result
}
