package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		// 去除重复数据
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			// 去除重复数据
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third > second {
				if nums[second]+nums[third] > target {
					third--
					continue
				} else {
					break
				}
			}
			if third == second {
				break
			}
			// 对结果进行审查
			if third > second && nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{},
			expected: [][]int{},
		},
		{
			input:    []int{1, 2, -2, -1},
			expected: [][]int{},
		},
		{
			input:    []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			input:    []int{-1, -1, -1, 2, 2, 2},
			expected: [][]int{{-1, -1, 2}},
		},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func main() {
	// 运行测试
	TestThreeSum(nil)
	fmt.Println("Tests finished.")
}
