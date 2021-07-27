package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: 2}
	node1 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 5}
	node3 := TreeNode{Val: 5}
	node4 := TreeNode{Val: 7}
	root.Left = &node1
	root.Right = &node2
	node2.Left = &node3
	node2.Right = &node4
	fmt.Println(findSecondMinimumValue(&root))
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	min2 := root.Val
	var bfs func(*TreeNode)
	bfs = func(node *TreeNode) {
		if node.Left == nil {
			return
		}
		if node.Left.Val > root.Val {
			if min2 == root.Val || min2 > node.Left.Val {
				min2 = node.Left.Val
				return
			}
		} else if node.Right.Val > root.Val {
			if min2 == root.Val || min2 > node.Right.Val {
				min2 = node.Right.Val
				return
			}
		} else {
			bfs(node.Left)
			bfs(node.Right)
		}
	}

	if min2 > root.Val {
		return min2
	} else {
		return -1
	}
}
