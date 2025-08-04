package LeetCode_Hot100_Round2

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	// 计算目标字符串数字量
	helpMapTarget := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		helpMapTarget[byte(p[i])] = helpMapTarget[byte(p[i])] + 1
	}
	// 初始数据
	result := make([]int, 0)
	for i := 0; i < len(p); i++ {
		if _, ok := helpMapTarget[byte(s[i])]; ok {
			helpMapTarget[byte(s[i])] = helpMapTarget[byte(s[i])] - 1
		}
	}
	flag := true
	for _, v := range helpMapTarget {
		if v != 0 {
			flag = false
			break
		}
	}
	if flag {
		result = append(result, 0)
	}

	// 滑动窗口统计
	for i := 0; i+len(p) < len(s); i++ {
		if _, ok := helpMapTarget[byte(s[i])]; ok {
			helpMapTarget[byte(s[i])] = helpMapTarget[byte(s[i])] + 1
		}
		if _, ok := helpMapTarget[byte(s[i+len(p)])]; ok {
			helpMapTarget[byte(s[i+len(p)])] = helpMapTarget[byte(s[i+len(p)])] - 1
		} else {
			continue
		}
		flag = true
		for _, v := range helpMapTarget {
			if v != 0 {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i+1)
		}
	}
	return result
}
