package main

func spiralOrder(matrix [][]int) []int {
	res := make([]int, len(matrix))
	i, j, xCnt, yCnt, cnt := 0, 0, len(matrix[0]), len(matrix), 0
	for {
		cnt += 1
		// 横向走
		if cnt%2 == 1 {
			// 增加方向走
			if (cnt/2)%2 == 0 {
				for j = cnt / 4; j < len(matrix[0])-cnt/4; j++ {
					res = append(res, matrix[i][j])
				}
			}
			// 减少方向走
			if (cnt/2)%2 == 1 {
				for i = len(matrix[0]) - cnt/4; i > cnt/4; i-- {
					res = append(res, matrix[i][j])
				}
			}
		}
		// 竖着走
		if cnt%2 == 0 {
			// 增加方向走
			if (cnt/2)%2 == 0 {
				for i = cnt / 4; i < len(matrix[0])-cnt/4; i++ {
					res = append(res, matrix[i][j])
				}
			}
			// 减少方向走
			if (cnt/2)%2 == 1 {
				for i = len(matrix[0]) - cnt/4; i > cnt/4; i-- {
					res = append(res, matrix[i][j])
				}
			}
		}
	}
}
