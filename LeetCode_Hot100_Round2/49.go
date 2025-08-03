package LeetCode_Hot100_Round2

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	helpMap := make(map[string][]string)
	for idx := 0; idx < len(strs); idx++ {
		s := []byte(strs[idx])
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		if value, ok := helpMap[string(s)]; ok {
			value = append(value, strs[idx])
			helpMap[string(s)] = value
		} else {
			helpMap[string(s)] = []string{strs[idx]}
		}
	}
	result := make([][]string, 0, len(helpMap))
	for _, v := range helpMap {
		result = append(result, v)
	}
	return result
}
