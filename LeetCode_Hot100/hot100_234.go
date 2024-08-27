package main

import (
	"fmt"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// Test cases for the isPalindrome function
func main() {
	// Test case 1: Empty list (should return true)
	var head1 *ListNode
	fmt.Println(isPalindrome(head1)) // true

	// Test case 2: Single element list (should return true)
	head2 := &ListNode{Val: 1}
	fmt.Println(isPalindrome(head2)) // true

	// Test case 3: Two elements, palindrome (should return true)
	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 1}}
	fmt.Println(isPalindrome(head3)) // true

	// Test case 4: Two elements, not palindrome (should return false)
	head4 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(isPalindrome(head4)) // false

	// Test case 5: Multiple elements, palindrome (should return true)
	head5 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	fmt.Println(isPalindrome(head5)) // true

	// Test case 6: Multiple elements, not palindrome (should return false)
	head6 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1}}}}
	fmt.Println(isPalindrome(head6)) // false
}

// Stub for isPalindrome function
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	} else {
		fast := head.Next
		slow := head
		// 快指针走到最后，慢指针走到一半
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
			if fast.Next == nil {
				break
			} else {
				fast = fast.Next
			}
		}
		nextAdd := slow.Next
		slow.Next = nil
		// 为链表后部颠倒顺序
		for nextAdd != nil {
			add := nextAdd
			nextAdd = add.Next
			temp := slow.Next
			slow.Next = add
			add.Next = temp
		}
		slow = slow.Next
		for slow != nil {
			if slow.Val != head.Val {
				return false
			}
			slow = slow.Next
			head = head.Next
		}
		return true
	}
}
