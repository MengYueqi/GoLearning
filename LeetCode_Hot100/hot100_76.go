package main

import "fmt"

func minWindow(s string, t string) string {
	cntMap := make(map[int]int)
	for i := 0; i < len(t); i++ {
		if _, exists := cntMap[int(t[i])]; exists == false {
			cntMap[int(t[i])] = 1
		} else {
			cntMap[int(t[i])] += 1
		}
	}
	fmt.Println(cntMap)
	first, last := 0, 0
	flag := false
	result := ""
	if _, exists := cntMap[int(s[last])]; exists {
		cntMap[int(s[last])] -= 1
	}
	flag = true
	for _, value := range cntMap {
		if value <= 0 {
			continue
		} else {
			flag = false
		}
	}
	if flag {
		result = s[first : last+1]
	}
	for last < len(s) {
		if flag == false {
			last += 1
			if last < len(s) {
				if _, exists := cntMap[int(s[last])]; exists {
					cntMap[int(s[last])] -= 1
					if cntMap[int(s[last])] == 0 {
						flag = true
						for _, value := range cntMap {
							if value <= 0 {
								continue
							} else {
								flag = false
							}
						}
					}
				}
				fmt.Println(cntMap)
				if flag {
					if len(result) == 0 {
						result = s[first : last+1]
					} else if len(s[first:last+1]) < len(result) {
						result = s[first : last+1]
					}
				}
			}
		} else {
			first += 1
			if _, exists := cntMap[int(s[first-1])]; exists {
				cntMap[int(s[first-1])] += 1
				if cntMap[int(s[first-1])] > 0 {
					flag = false
				}
			}
			if flag {
				if len(result) == 0 {
					result = s[first : last+1]
				} else if len(s[first:last+1]) < len(result) {
					result = s[first : last+1]
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(minWindow("a", "aa"))
}
