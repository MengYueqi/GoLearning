package main

import (
	"fmt"
	"strings"
)

var solution [][]string

func backtrack(queens []int, curRow, n int, col, firstTag, secondTag map[int]bool) {
	if curRow == n {
		var result []string
		for i := 0; i < n; i++ {
			var builder strings.Builder
			for j := 0; j < n; j++ {
				if j == queens[i] {
					builder.WriteString("Q")
				} else {
					builder.WriteString(".")
				}
			}
			result = append(result, builder.String())
		}
		solution = append(solution, result)
	} else {
		// 轮询所有的可能列
		for i := 0; i < n; i++ {
			if !firstTag[i+curRow] && !secondTag[i-curRow] && !col[i] {
				// 可以进行添加
				firstTag[i+curRow] = true
				secondTag[i-curRow] = true
				col[i] = true
				queens[curRow] = i
				backtrack(queens, curRow+1, n, col, firstTag, secondTag)
				firstTag[i+curRow] = false
				secondTag[i-curRow] = false
				col[i] = false
			}
		}
	}
}

func solveNQueens(n int) [][]string {
	queens := make([]int, n)
	backtrack(queens, 0, n, make(map[int]bool), make(map[int]bool), make(map[int]bool))
	return solution
}

func main() {
	fmt.Println(solveNQueens(1))
}
