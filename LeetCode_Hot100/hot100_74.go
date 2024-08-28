package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) / 2
	left, right := 0, len(matrix)-1
	for left <= right {
		if matrix[i][0] <= target && matrix[i][len(matrix[0])-1] >= target {
			break
		} else {
			if matrix[i][0] > target {
				// 在 i-1 和左边界之间查找
				right = i - 1
				i = (left + right) / 2
			} else {
				left = i + 1
				i = (right + left) / 2
			}
		}
	}
	if left > right {
		return false
	}
	left, right = 0, len(matrix[0])-1
	j := len(matrix[0]) / 2
	for left <= right {
		fmt.Println(i, j)
		if matrix[i][j] == target {
			return true
		} else {
			if matrix[i][j] > target {
				right = j - 1
				j = (left + right) / 2
			} else {
				left = j + 1
				j = (left + right) / 2
			}
		}
	}
	return false
}

func main() {
	searchMatrix([][]int{{1}, {3}, {5}}, 1)
}
