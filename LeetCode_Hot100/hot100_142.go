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
func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]int)
	if head == nil {
		return nil
	}
	for head.Next != nil {
		mp[head] = 1
		_, exists := mp[head.Next]
		if exists {
			return head.Next
		}
		head = head.Next
	}
	return nil
}

func main() {
	// Test case 1: Empty list
	var head1 *ListNode
	fmt.Println("Test case 1 (Empty list):", detectCycle(head1)) // Expected output: <nil>

	// Test case 2: Single node without cycle
	head2 := &ListNode{Val: 1}
	fmt.Println("Test case 2 (Single node without cycle):", detectCycle(head2)) // Expected output: <nil>

	// Test case 3: Single node with cycle
	head3 := &ListNode{Val: 1}
	head3.Next = head3
	fmt.Println("Test case 3 (Single node with cycle):", detectCycle(head3)) // Expected output: Node with Val=1

	// Test case 4: Multiple nodes without cycle
	head4 := &ListNode{Val: 1}
	head4.Next = &ListNode{Val: 2}
	head4.Next.Next = &ListNode{Val: 3}
	head4.Next.Next.Next = &ListNode{Val: 4}
	fmt.Println("Test case 4 (Multiple nodes without cycle):", detectCycle(head4)) // Expected output: <nil>

	// Test case 5: Multiple nodes with cycle
	head5 := &ListNode{Val: 1}
	head5.Next = &ListNode{Val: 2}
	head5.Next.Next = &ListNode{Val: 3}
	head5.Next.Next.Next = &ListNode{Val: 4}
	head5.Next.Next.Next.Next = head5.Next                                      // Create cycle
	fmt.Println("Test case 5 (Multiple nodes with cycle):", detectCycle(head5)) // Expected output: Node with Val=2

	// Test case 6: Longer list without cycle
	head6 := &ListNode{Val: 1}
	curr := head6
	for i := 2; i <= 10; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	fmt.Println("Test case 6 (Longer list without cycle):", detectCycle(head6)) // Expected output: <nil>

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
	curr.Next = cycleStart                                                   // Create cycle
	fmt.Println("Test case 7 (Longer list with cycle):", detectCycle(head7)) // Expected output: Node with Val=5
}
