package main

import "fmt"

func partitionLabels(s string) []int {
	farthestChar := make(map[int]int)
	// 统计最远字符位置
	for i := 0; i < len(s); i++ {
		_, exist := farthestChar[int(s[i])]
		if exist {
			farthestChar[int(s[i])] = i
		} else {
			farthestChar[int(s[i])] = i
		}
	}
	// 获取字符串分组
	var result []int
	for i := 0; i < len(s); i++ {
		first := i
		lastChar := farthestChar[int(s[i])]
		for i < lastChar {
			lastChar = max(lastChar, farthestChar[int(s[i])])
			i++
		}
		result = append(result, i-first+1)
	}
	return result
}

func main() {
	fmt.Print(partitionLabels("ababcbacadefegdehijhklij"))
}
