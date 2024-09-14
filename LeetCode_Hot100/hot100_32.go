package main

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < len(dp); i++ {
		if s[i-1] == ')' {
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else if i-1-dp[i-1]-1 >= 0 && s[i-1-dp[i-1]-1] == '(' {
				dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]-2]
			} else {
				dp[i] = 0
			}
		} else {
			dp[i] = 0
		}
	}
	maxNum := 0
	for i := 0; i < len(dp); i++ {
		maxNum = max(maxNum, dp[i])
	}
	return maxNum
}
