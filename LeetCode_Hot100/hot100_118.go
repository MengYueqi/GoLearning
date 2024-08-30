package main

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 1 {
		return [][]int{{1}}
	} else {
		result = append(result, []int{1})
		for i := 1; i < numRows; i++ {
			var line []int
			line = append(line, 1)
			for j := 1; j < i; j++ {
				line = append(line, result[i-1][j-1]+result[i-1][j])
			}
			line = append(line, 1)
			result = append(result, line)
		}
	}
	return result
}
