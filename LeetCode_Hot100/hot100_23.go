package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	flag := 0
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		} else {
			flag = 1
		}
	}
	if flag == 0 {
		return nil
	}
	// 获取首个数据
	var result *ListNode
	minIndex := -1
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			if minIndex == -1 {
				minIndex = i
			} else if lists[i].Val < lists[minIndex].Val {
				minIndex = i
			}
		}
	}
	result = lists[minIndex]
	lists[minIndex] = lists[minIndex].Next
	temp := result

	for {
		minIndex = -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if minIndex == -1 {
					minIndex = i
				} else if lists[i].Val < lists[minIndex].Val {
					minIndex = i
				}
			}
		}
		// 根据 minIndex 获取是否退出
		if minIndex == -1 {
			break
		}
		temp.Next = lists[minIndex]
		temp = temp.Next
		lists[minIndex] = lists[minIndex].Next
	}
	return result
}
