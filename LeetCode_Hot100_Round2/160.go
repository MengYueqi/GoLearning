package LeetCode_Hot100_Round2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 构建 headA 的链表的哈希表
	hm := make(map[*ListNode]struct{})
	for h := headA; h != nil; h = h.Next {
		hm[h] = struct{}{}
	}

	// 在 headA 哈希表中查找是否有重合
	for h := headB; h != nil; h = h.Next {
		if _, ok := hm[h]; ok {
			return h
		}
	}
	return nil
}
