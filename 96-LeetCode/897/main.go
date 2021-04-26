package main

func main() {
	var ret *TreeNode = increasingBST(new(TreeNode))
}

/**
 * Definition for a binary tree node.
 *
 *
 *
 *
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	retNode := &TreeNode{}
	resNode := retNode

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		// 在中序遍历中修改节点**
		resNode.Right = node
		node.Left = nil
		resNode = node

		inOrder(node.Right)
	}
	inOrder(root)
	return retNode
}
