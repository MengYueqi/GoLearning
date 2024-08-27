package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, result := depth(root)
	return result
}

func depth(root *TreeNode) (d int, maxL int) {
	if root == nil {
		return -1, 0
	}
	dLeft, maxLeft := depth(root.Left)
	dRight, maxRight := depth(root.Right)
	return max(dLeft, dRight) + 1, max(maxLeft, maxRight, dLeft+dRight+2)
}

func main() {

}
