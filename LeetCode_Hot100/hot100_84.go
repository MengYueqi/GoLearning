package main

import "fmt"

type heightWithIdx struct {
	height int
	idx    int
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	heightStack := make([]heightWithIdx, 0)
	heightStack = append(heightStack, heightWithIdx{height: 0, idx: -1})
	first := make([]int, len(heights))
	last := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		// 如果比栈顶大就进栈
		if heights[i] > heightStack[len(heightStack)-1].height {
			heightStack = append(heightStack, heightWithIdx{height: heights[i], idx: i})
		} else {
			//	维护单调增长栈 将相关元素出栈并记录前后 idx
			for j := len(heightStack) - 1; j > 0; j-- {
				if heights[i] < heightStack[j].height {
					first[heightStack[j].idx] = heightStack[j-1].idx
					last[heightStack[j].idx] = i
					heightStack = heightStack[:len(heightStack)-1]
				} else {
					break
				}
			}
			heightStack = append(heightStack, heightWithIdx{height: heights[i], idx: i})
		}
		fmt.Println(heightStack)
	}

	fmt.Println(first)
	fmt.Println(last)
	result := 0
	for i := 0; i < len(heights); i++ {
		result = max(heights[i]*(last[i]-first[i]-1), result)
	}
	return result
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
