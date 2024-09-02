package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = false
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if len(wordDict[j]) == i+1 {
				if wordDict[j] == s[:i+1] {
					dp[i] = true
				}
			} else {
				if len(wordDict[j]) < i+1 {
					if dp[i-len(wordDict[j])] && wordDict[j] == s[i-len(wordDict[j])+1:i+1] {
						dp[i] = true
					}
				}
			}
		}
	}
	return dp[len(s)-1]
}
