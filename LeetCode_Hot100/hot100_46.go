package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backtracking(nums, 0)
	return res
}

func backtracking(nums []int, first int) {
	if first == len(nums)-1 {
		perm := make([]int, len(nums))
		copy(perm, nums)
		res = append(res, perm)
		return
	}
	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtracking(nums, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
	return
}

func main() {
	// 测试用例1：空数组
	nums1 := []int{}
	result1 := permute(nums1)
	fmt.Println("测试用例1：", result1)

	// 测试用例2：单个元素的数组
	nums2 := []int{1}
	result2 := permute(nums2)
	fmt.Println("测试用例2：", result2)

	// 测试用例3：两个元素的数组
	nums3 := []int{1, 2}
	result3 := permute(nums3)
	fmt.Println("测试用例3：", result3)

	// 测试用例4：三个元素的数组
	nums4 := []int{1, 2, 3}
	result4 := permute(nums4)
	fmt.Println("测试用例4：", result4)

	// 测试用例5：四个元素的数组
	nums5 := []int{1, 2, 3, 4}
	result5 := permute(nums5)
	fmt.Println("测试用例5：", result5)
}
