package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 先根据每个子数组的第一个元素排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)
	result := make([][]int, 0)

	first := -1
	last := -1
	for _, interval := range intervals {
		println(interval[0], interval[1])
		if first == -1 {
			first = interval[0]
			last = interval[1]
		} else {
			if last >= interval[0] {
				last = max(interval[1], last)
			} else {
				// 将数据添加到结果中
				result = append(result, []int{first, last})
				first = interval[0]
				last = interval[1]
			}
		}
	}
	result = append(result, []int{first, last})

	return result
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {8, 10}, {5, 6}, {8, 18}}))
}
