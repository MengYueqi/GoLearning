package LeetCode_Hot100_Round2

func longestConsecutive(nums []int) int {
	// 对空数组处理
	if len(nums) == 0 {
		return 0
	}

	// 获取最大最小值
	minValue, maxValue := nums[0], nums[0]
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
		if num < minValue {
			minValue = num
		}
	}

	// 将所有数据填入辅助 hashMap
	hm := make(map[int]struct{})
	for _, num := range nums {
		hm[num] = struct{}{}
	}

	// 获取最长序列
	maxLen, curLen := 1, 1
	for k, _ := range hm {
		// 只在连续序列的首个才进行遍历
		if _, ok := hm[k-1]; ok {
			continue
		} else {
			// 连续数量统计
			next := k + 1
			for ; next <= maxValue; next++ {
				if _, ok := hm[next]; ok {
					curLen += 1
				} else {
					break
				}
			}
			maxLen = max(maxLen, curLen)
			curLen = 1
		}
	}
	return maxLen
}
