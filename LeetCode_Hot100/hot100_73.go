package main

import "fmt"

func setZeroes(matrix [][]int) {
	set0Is := make(map[int]int)
	set0Js := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				set0Is[i] = 1
				set0Js[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, existsI := set0Is[i]
			_, existsJ := set0Js[j]
			if existsI || existsJ {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	// 测试用例是GPT写的，有错误
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			output: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			output: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			input: [][]int{
				{1, 0, 3},
				{4, 5, 6},
				{0, 8, 9},
			},
			output: [][]int{
				{0, 0, 0},
				{0, 5, 6},
				{0, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		setZeroes(testCase.input)
		fmt.Printf("Input: %v\n", testCase.input)
		fmt.Printf("Expected Output: %v\n\n", testCase.output)
	}
}
