package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var ret *ListNode
	ret = nil
	for head != nil {
		if ret == nil {
			tempNode := new(ListNode)
			tempNode.Val = head.Val
			tempNode.Next = nil
			ret = tempNode
		} else {
			tempNode := new(ListNode)
			tempNode.Val = head.Val
			temp := ret
			ret = tempNode
			ret.Next = temp
		}
		head = head.Next
	}
	return ret
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original list:")
	printList(head)

	// 反转链表
	reversedHead := reverseList(head)

	fmt.Println("Reversed list:")
	printList(reversedHead)
}
