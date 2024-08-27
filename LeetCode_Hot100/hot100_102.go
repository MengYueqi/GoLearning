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
func levelOrder(root *TreeNode) [][]int {
	// 对空树进行判断
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		cnt := len(queue)
		levelRes := make([]int, 0)
		for i := 0; i < cnt; i++ {
			levelRes = append(levelRes, queue[0].Val)
			// 将孩子节点加入队列并将父节点出队
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		result = append(result, levelRes)
	}
	return result
}

func main() {
	// 测试用例1：普通二叉树
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 7}

	fmt.Println("测试用例1：", levelOrder(root1)) // 输出: [[1] [2 3] [4 5 6 7]]

	// 测试用例2：只有一个节点
	root2 := &TreeNode{Val: 1}
	fmt.Println("测试用例2：", levelOrder(root2)) // 输出: [[1]]

	// 测试用例3：空树
	var root3 *TreeNode
	fmt.Println("测试用例3：", levelOrder(root3)) // 输出: []

	// 测试用例4：不平衡二叉树
	root4 := &TreeNode{Val: 1}
	root4.Left = &TreeNode{Val: 2}
	root4.Left.Left = &TreeNode{Val: 3}
	root4.Left.Left.Left = &TreeNode{Val: 4}

	fmt.Println("测试用例4：", levelOrder(root4)) // 输出: [[1] [2] [3] [4]]

	// 测试用例5：另一棵普通二叉树
	root5 := &TreeNode{Val: 1}
	root5.Left = &TreeNode{Val: 2}
	root5.Right = &TreeNode{Val: 3}
	root5.Left.Right = &TreeNode{Val: 4}
	root5.Right.Left = &TreeNode{Val: 5}

	fmt.Println("测试用例5：", levelOrder(root5)) // 输出: [[1] [2 3] [4 5]]
}
