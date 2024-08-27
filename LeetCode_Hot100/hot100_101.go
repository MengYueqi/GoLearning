package main

import "fmt"

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return check(root.Left, root.Right)
	}
}

func check(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && check(root1.Left, root2.Right) && check(root1.Right, root2.Left)
}

func main() {
	// Helper function to create a binary tree from a slice
	createBinaryTree := func(vals []interface{}) *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		nodes := make([]*TreeNode, len(vals))
		for i, val := range vals {
			if val != nil {
				nodes[i] = &TreeNode{Val: val.(int)}
			}
		}
		for i := 0; i < len(vals); i++ {
			if nodes[i] != nil {
				if 2*i+1 < len(vals) {
					nodes[i].Left = nodes[2*i+1]
				}
				if 2*i+2 < len(vals) {
					nodes[i].Right = nodes[2*i+2]
				}
			}
		}
		return nodes[0]
	}

	// Test case 1: Symmetric tree
	tree1 := createBinaryTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	fmt.Println("Test Case 1:", isSymmetric(tree1)) // Expected: true

	// Test case 2: Asymmetric tree
	tree2 := createBinaryTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	fmt.Println("Test Case 2:", isSymmetric(tree2)) // Expected: false

	// Test case 3: Single node tree
	tree3 := createBinaryTree([]interface{}{1})
	fmt.Println("Test Case 3:", isSymmetric(tree3)) // Expected: true

	// Test case 4: Empty tree
	var tree4 *TreeNode
	fmt.Println("Test Case 4:", isSymmetric(tree4)) // Expected: true

	// Test case 5: Larger symmetric tree
	tree5 := createBinaryTree([]interface{}{1, 2, 2, 3, 4, 4, 3, 5, 6, 6, 5, 5, 6, 6, 5})
	fmt.Println("Test Case 5:", isSymmetric(tree5)) // Expected: true

	// Test case 6: Larger asymmetric tree
	tree6 := createBinaryTree([]interface{}{1, 2, 2, 3, 4, 4, 3, 5, nil, nil, 5})
	fmt.Println("Test Case 6:", isSymmetric(tree6)) // Expected: false
}
