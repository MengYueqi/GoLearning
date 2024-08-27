package main

import (
	"container/heap"
	"fmt"
)

var a []int

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool  { return a[h[i]] > a[h[j]] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := make(hp, k)
	for i := 0; i < k; i++ {
		q[i] = i
	}
	heap.Init(&q)

	var res []int
	res = append(res, a[q[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(&q, i)
		for q[0] <= i-k {
			heap.Pop(&q)
		}
		res = append(res, a[q[0]])
	}
	return res
}

func main() {
	// 测试用例 1
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Println("测试用例 1:", maxSlidingWindow(nums1, k1)) // 预期输出: [3 3 5 5 6 7]

	// 测试用例 2
	nums2 := []int{9, 11, 8, 5, 7, 10}
	k2 := 2
	fmt.Println("测试用例 2:", maxSlidingWindow(nums2, k2)) // 预期输出: [11 11 8 7 10]

	// 测试用例 3
	nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k3 := 4
	fmt.Println("测试用例 3:", maxSlidingWindow(nums3, k3)) // 预期输出: [4 5 6 7 8 9 10]

	// 测试用例 4
	nums4 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	k4 := 5
	fmt.Println("测试用例 4:", maxSlidingWindow(nums4, k4)) // 预期输出: [10 9 8 7 6]

	// 测试用例 5
	nums5 := []int{4, 3, 5, 4, 3, 3, 6, 7}
	k5 := 3
	fmt.Println("测试用例 5:", maxSlidingWindow(nums5, k5)) // 预期输出: [5 5 5 4 6 7]

	// 测试用例 5
	nums6 := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	k6 := 7
	fmt.Println("测试用例 5:", maxSlidingWindow(nums6, k6)) // 预期输出: [5 5 5 4 6 7]
}
