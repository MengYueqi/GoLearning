package LeetCode_Hot100_Round2

func setZeroes(matrix [][]int) {
	// 构建记录要置为 0 的行列集合
	l, r := make(map[int]struct{}), make(map[int]struct{})

	// 记录所有为 0 的行和列
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				l[i] = struct{}{}
				r[j] = struct{}{}
			}
		}
	}

	// 对矩阵进行置 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if _, ok := l[i]; ok {
				matrix[i][j] = 0
			} else if _, ok := r[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

}
