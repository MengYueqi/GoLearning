package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root == p || root == q {
		return root
	}
	leftNode, rightNode := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if leftNode == nil {
		return rightNode
	} else if rightNode == nil {
		return leftNode
	} else {
		return root
	}
}
