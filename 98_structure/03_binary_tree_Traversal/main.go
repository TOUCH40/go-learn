package main

// tree struct
type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {

}

// create a binary tree
func CreateTreeFromIntArray(ar []int) *TreeNode {
	queue:=[]TreeNode
	root:=new(TreeNode)
	queue = append(queue,root)
	for _,item in range ar{
		tem:=new(TreeNode)
		tem.value = item
		tem=append(tem)
	}
}
