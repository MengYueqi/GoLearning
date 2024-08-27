package main

import (
	"fmt"
	"reflect"
)

func rotate(nums []int, k int) {
	k = k % len(nums)
	result := nums[len(nums)-k:]
	result = append(result, nums[0:len(nums)-k]...)
	copy(nums, result)
	return
}

func testRotate() {
	tests := []struct {
		input  []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 10, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, test := range tests {
		rotate(test.input, test.k)
		if !reflect.DeepEqual(test.input, test.expect) {
			fmt.Printf("rotate(%v, %d) = %v; expect %v\n", test.input, test.k, test.input, test.expect)
		} else {
			fmt.Printf("rotate(%v, %d) passed\n", test.input, test.k)
		}
	}
}

func main() {
	testRotate()
}
