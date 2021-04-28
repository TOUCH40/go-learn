package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {

	sum := 0
	preOrder := func(t *TreeNode) {
		if t.Val >= low && t.Val <= high {
			sum += t.Val
		}
		preOrder(t.Left)
		preOrder(t.Right)
	}
	return sum
}
func main() {

}
