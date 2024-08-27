package main

import "fmt"

func subarraySum(nums []int, k int) int {
	cnt, preSum := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			cnt += m[preSum-k]
		}
		m[preSum] += 1
	}
	return cnt
}

func main() {
	// 定义测试用例
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 1, 1}, k: 2, want: 2},
		{nums: []int{1, 2, 3}, k: 3, want: 2},
		{nums: []int{-1, -1, 1}, k: 0, want: 1},
		{nums: []int{1, 2, 1, 2, 1}, k: 3, want: 4},
		{nums: []int{1, -1, 0}, k: 0, want: 3},
	}

	for _, tc := range testCases {
		got := subarraySum(tc.nums, tc.k)
		if got != tc.want {
			fmt.Printf("For nums=%v, k=%d: expected %d, but got %d\n", tc.nums, tc.k, tc.want, got)
		} else {
			fmt.Printf("Test passed for nums=%v, k=%d: got %d\n", tc.nums, tc.k, got)
		}
	}

}
