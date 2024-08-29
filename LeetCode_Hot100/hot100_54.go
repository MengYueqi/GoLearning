package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// 设置方向：1 向右，2 向下，3 向左，4 向上
	directions := 1
	cnt := make([][]int, len(matrix))
	for i := range cnt {
		cnt[i] = make([]int, len(matrix[0]))
	}
	i, j := 0, 0
	var result []int
	flag := 1
	first := 1
	for flag == 1 {
		flag = 0
		if directions == 1 {
			if first == 1 {
				first = 0
			} else {
				j++
			}
			for i <= len(matrix)-1 && i >= 0 && j < len(matrix[0]) && j >= 0 && cnt[i][j] == 0 {
				cnt[i][j] += 1
				result = append(result, matrix[i][j])
				j++
				flag = 1
			}
			j--
		} else if directions == 2 {
			i++
			for i <= len(matrix)-1 && i >= 0 && j < len(matrix[0]) && j >= 0 && cnt[i][j] == 0 {
				cnt[i][j] += 1
				result = append(result, matrix[i][j])
				i++
				flag = 1
			}
			i--
		} else if directions == 3 {
			j--
			for i <= len(matrix)-1 && i >= 0 && j < len(matrix[0]) && j >= 0 && cnt[i][j] == 0 {
				cnt[i][j] += 1
				result = append(result, matrix[i][j])
				j--
				flag = 1
			}
			j++
		} else if directions == 4 {
			i--
			for i <= len(matrix)-1 && i >= 0 && j < len(matrix[0]) && j >= 0 && cnt[i][j] == 0 {
				cnt[i][j] += 1
				result = append(result, matrix[i][j])
				i--
				flag = 1
			}
			i++
		}

		if directions == 4 {
			directions = 1
		} else {
			directions++
		}
	}
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
