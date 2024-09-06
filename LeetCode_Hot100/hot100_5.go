package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			dp[i][j] = false
		}
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j+i < len(s); j++ {
			if i == 1 {
				dp[j][j+i] = s[j] == s[j+i]
			} else {
				dp[j][j+i] = dp[j+1][j+i-1] && s[j] == s[j+i]
			}
		}
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(dp[i])
	}
	for i := len(s); i > 0; i-- {
		for j := 0; j+i < len(s); j++ {
			if dp[j][j+i] {
				return s[j : j+i+1]
			}
		}
	}
	return s[0:1]
}

func main() {
	longestPalindrome("abb")
}
