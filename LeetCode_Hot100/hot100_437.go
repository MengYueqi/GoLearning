package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rootSum(root *TreeNode, targetSum int, curSum int) int {
	if root == nil {
		return 0
	}
	if curSum+root.Val == targetSum {
		return rootSum(root.Left, targetSum, curSum+root.Val) + rootSum(root.Right, targetSum, curSum+root.Val) + 1
	} else {
		return rootSum(root.Left, targetSum, curSum+root.Val) + rootSum(root.Right, targetSum, curSum+root.Val)
	}
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + rootSum(root, targetSum, 0)
}
