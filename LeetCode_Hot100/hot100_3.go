package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	lenOfLS := 0
	for right < len(s) {
		if strings.Contains(s[left:right], string(s[right])) {
			left = left + 1
		} else {
			right = right + 1
		}

		if lenOfLS < right-left {
			lenOfLS = right - left
		}
	}
	return lenOfLS
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 输出 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 输出 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 输出 3
	fmt.Println(lengthOfLongestSubstring(""))         // 输出 0
	fmt.Println(lengthOfLongestSubstring(" "))        // 输出 1
	fmt.Println(lengthOfLongestSubstring("au"))       // 输出 2
	fmt.Println(lengthOfLongestSubstring("aab"))      // 输出 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 输出 3
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))  // 输出 5
	fmt.Println(lengthOfLongestSubstring("cdd"))      // 输出 2
	fmt.Println(lengthOfLongestSubstring("abba"))     // 输出 2
}
