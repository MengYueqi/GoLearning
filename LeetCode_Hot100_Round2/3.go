package LeetCode_Hot100_Round2

// package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	hm := make(map[byte]int)
	maxLen := 0
	first, last := 0, 0
	for ; last < len(s); last++ {
		if value, ok := hm[byte(s[last])]; ok && value != 0 {
			hm[byte(s[last])] = hm[byte(s[last])] + 1
			// 向前移动 first 指针到无重复子串
			for ; ; first++ {
				hm[byte(s[first])] = hm[byte(s[first])] - 1
				if s[first] == s[last] {
					first = first + 1
					break
				}
			}
			maxLen = max(maxLen, last-first+1)
		} else {
			// 滑动窗口不包含此字符
			hm[byte(s[last])] = 1
			maxLen = max(maxLen, last-first+1)
		}
		fmt.Println(maxLen)
	}
	return maxLen
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
}
