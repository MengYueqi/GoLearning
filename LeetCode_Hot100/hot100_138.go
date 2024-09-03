package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	newHead := &Node{Val: head.Val, Next: nil, Random: nil}
	nodeMap[head] = newHead
	source, result := head, newHead
	for source.Next != nil {
		temp := &Node{Val: source.Next.Val, Next: nil, Random: nil}
		result.Next = temp
		nodeMap[source.Next] = temp
		result, source = result.Next, source.Next
	}
	source, result = head, newHead
	for source != nil {
		result.Random = nodeMap[source.Random]
		result, source = result.Next, source.Next
	}
	return newHead
}
