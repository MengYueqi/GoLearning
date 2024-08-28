package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left, right := root.Left, root.Right
	root.Left = nil
	flatten(left)
	flatten(right)
	root.Right = left
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	temp.Right = right
}
