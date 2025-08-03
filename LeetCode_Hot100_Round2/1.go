package LeetCode_Hot100_Round2

func twoSum(nums []int, target int) []int {
	helpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := helpMap[nums[i]]; ok {
			return []int{i, value}
		} else {
			helpMap[target-nums[i]] = i
		}
	}
	return []int{}
}
