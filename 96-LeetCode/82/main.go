package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// num := []int{1, 2, 3, 4}
	nb := deleteDuplicates(num)
	fmt.Printf("%b", nb)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var fnode = head
	for fnode.Next != nil && fnode.Next.Next != nil {
		if fnode.Val == fnode.Next.Val {
			tnode = fnode.Next
			fnode.Next = fnode.Next.Next
			tnode = nil
		} else {
			fnode = fnode.Next
		}
	}
}
