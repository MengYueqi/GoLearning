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
func hasCycle(head *ListNode) bool {
	// 使用快慢指针
	fast, slow := head, head

	if head == nil {
		return false
	}

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1: Empty list
	var head1 *ListNode
	fmt.Println("Test case 1 (Empty list):", hasCycle(head1)) // Expected output: false

	// Test case 2: Single node without cycle
	head2 := &ListNode{Val: 1}
	fmt.Println("Test case 2 (Single node without cycle):", hasCycle(head2)) // Expected output: false

	// Test case 3: Single node with cycle
	head3 := &ListNode{Val: 1}
	head3.Next = head3
	fmt.Println("Test case 3 (Single node with cycle):", hasCycle(head3)) // Expected output: true

	// Test case 4: Multiple nodes without cycle
	head4 := &ListNode{Val: 1}
	head4.Next = &ListNode{Val: 2}
	head4.Next.Next = &ListNode{Val: 3}
	head4.Next.Next.Next = &ListNode{Val: 4}
	fmt.Println("Test case 4 (Multiple nodes without cycle):", hasCycle(head4)) // Expected output: false

	// Test case 5: Multiple nodes with cycle
	head5 := &ListNode{Val: 1}
	head5.Next = &ListNode{Val: 2}
	head5.Next.Next = &ListNode{Val: 3}
	head5.Next.Next.Next = &ListNode{Val: 4}
	head5.Next.Next.Next.Next = head5.Next                                   // Create cycle
	fmt.Println("Test case 5 (Multiple nodes with cycle):", hasCycle(head5)) // Expected output: true

	// Test case 6: Longer list without cycle
	head6 := &ListNode{Val: 1}
	curr := head6
	for i := 2; i <= 10; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	fmt.Println("Test case 6 (Longer list without cycle):", hasCycle(head6)) // Expected output: false

	// Test case 7: Longer list with cycle
	head7 := &ListNode{Val: 1}
	curr = head7
	var cycleStart *ListNode
	for i := 2; i <= 10; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
		if i == 5 {
			cycleStart = curr
		}
	}
	curr.Next = cycleStart                                                // Create cycle
	fmt.Println("Test case 7 (Longer list with cycle):", hasCycle(head7)) // Expected output: true
}
