package main

import (
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	valNodeMap := make(map[int][]*ListNode)
	temp := head
	for temp != nil {
		nodeList, exists := valNodeMap[temp.Val]
		if !exists {
			valNodeMap[temp.Val] = []*ListNode{temp}
			temp = temp.Next
		} else {
			valNodeMap[temp.Val] = append(nodeList, temp)
			temp = temp.Next
		}
	}
	var values []int
	for key := range valNodeMap {
		values = append(values, key)
	}
	sort.Ints(values)
	result := valNodeMap[values[0]][0]
	temp = result
	for i := 1; i < len(valNodeMap[values[0]]); i++ {
		temp.Next = valNodeMap[values[0]][i]
		temp = temp.Next
	}
	for i := 1; i < len(values); i++ {
		nodeList := valNodeMap[values[i]]
		for j := 0; j < len(nodeList); j++ {
			temp.Next = valNodeMap[values[i]][j]
			temp = temp.Next
		}
	}
	temp.Next = nil
	return result
}
