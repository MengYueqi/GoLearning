package LeetCode_Hot100_Round2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
	}
}
