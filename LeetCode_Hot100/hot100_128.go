package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	maxCnt, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			cnt = 1
		} else if nums[i] == nums[i-1]+1 {
			cnt += 1
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			cnt = 1
		}
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}

func main() {
	fmt.Println(longestConsecutive([]int{1, 9, 6, 10, 3, 2}))     // 3
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))  // 4
	fmt.Println(longestConsecutive([]int{0, -1, -2, 2, 3, 4, 1})) // 5
	fmt.Println(longestConsecutive([]int{}))                      // 0
	fmt.Println(longestConsecutive([]int{5, 5, 5, 5}))            // 1
	fmt.Println(longestConsecutive([]int{10, 5, 12, 11, 6, 7}))   // 4
	fmt.Println(longestConsecutive([]int{1, 2, 3, 4, 5}))         // 5
}
