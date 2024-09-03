package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum = 0

func contributeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCon, rightCon := max(contributeNode(root.Left), 0), max(contributeNode(root.Right), 0)
	sum = max(leftCon+root.Val+rightCon, sum)
	return max(leftCon, rightCon) + root.Val
}

func maxPathSum(root *TreeNode) int {
	sum = root.Val
	contributeNode(root)
	return sum
}
