package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1

	for x < len(matrix) && y >= 0 {
		if target == matrix[x][y] {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
}
