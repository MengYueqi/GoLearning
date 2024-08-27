package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	result := temp
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			fmt.Println(list1.Val, list2.Val)
			if list1.Val <= list2.Val {
				temp.Next = list1
				temp = temp.Next
				list1 = list1.Next
			} else {
				temp.Next = list2
				temp = temp.Next
				list2 = list2.Next
			}
		} else {
			if list1 != nil {
				temp.Next = list1
				temp = temp.Next
				list1 = list1.Next
			}
			if list2 != nil {
				temp.Next = list2
				temp = temp.Next
				list2 = list2.Next
			}
		}
	}
	return result.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// Helper function to create a linked list from a slice
	createLinkedList := func(vals []int) *ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &ListNode{Val: vals[0]}
		current := head
		for _, val := range vals[1:] {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		return head
	}

	// Test case 1: Both lists are empty
	var list1, list2 *ListNode
	result := mergeTwoLists(list1, list2)
	fmt.Print("Test Case 1: ")
	printList(result)

	// Test case 2: First list is empty, second list is not
	list1 = createLinkedList([]int{})
	list2 = createLinkedList([]int{1, 3, 5})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Test Case 2: ")
	printList(result)

	// Test case 3: First list is not empty, second list is empty
	list1 = createLinkedList([]int{2, 4, 6})
	list2 = createLinkedList([]int{})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Test Case 3: ")
	printList(result)

	// Test case 4: Both lists have elements
	list1 = createLinkedList([]int{1, 2, 4})
	list2 = createLinkedList([]int{1, 3, 4})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Test Case 4: ")
	printList(result)

	// Test case 5: Lists with different lengths
	list1 = createLinkedList([]int{1, 3, 5, 7})
	list2 = createLinkedList([]int{2, 4, 6})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Test Case 5: ")
	printList(result)
}
