package main

 import("fmt")

func main(){
v1:=new(TreeNode)
v1.Val=1
v2:=new(TreeNode)
v2.Val=2
var bl bool =leafSimilar(v1,v2)
fmt.Printf("%b",bl)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    nodes1:=[]int{}
    nodes2:=[]int{}
    preOrder(root1,&nodes1)
    preOrder(root2,&nodes2)
    if(len(nodes1)!=len(nodes2)){
        return false
    }
    for _i:=range nodes1{
        if(nodes1[_i]!=nodes2[_i]){
            return false
        }
    }
    return true
}

func preOrder(r *TreeNode,v *[]int){
    if r==nil{
        return
    }
    if(r.Left==nil&&r.Right==nil){
        *v=append(*v,r.Val)
    }
    preOrder(r.Left,v)
    preOrder(r.Right,v)
}