package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, make([]int, 0))
	for i := 0; i < len(nums); i++ {
		addSets := make([][]int, 0)
		for j := 0; j < len(res); j++ {
			as := make([]int, len(res[j])+1)
			copy(as, append(res[j], nums[i]))
			addSets = append(addSets, as)
			fmt.Println(res[j], nums[i])
		}
		res = append(res, addSets...)
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
