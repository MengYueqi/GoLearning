package main

import "fmt"

func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	// 初始化 dp 数组
	for i := 1; i < len(s); i++ {
		for j := 0; j+i < len(s); j++ {
			if i == 1 {
				dp[j][j+i] = s[j] == s[j+i]
			} else {
				dp[j][j+i] = dp[j+1][j+i-1] && (s[j] == s[j+i])
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}

	var dfs func(idx int, stringSet []string)
	result := make([][]string, 0)
	dfs = func(idx int, stringSet []string) {
		if idx == len(s) {
			// 这里要做一个拷贝，避免共享底层数组
			copySet := make([]string, len(stringSet))
			copy(copySet, stringSet)
			result = append(result, copySet)
			return
		}
		for i := 0; i+idx < len(s); i++ {
			if dp[idx][i+idx] == true {
				//fmt.Println(append(stringSet, s[idx:i+idx+1]))
				dfs(idx+i+1, append(stringSet, s[idx:i+idx+1]))
			}
		}
	}
	dfs(0, make([]string, 0))
	return result
}

func main() {
	result := partition("ababbbabbaba")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
