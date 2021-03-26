/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	Val int
	Next *ListNode
}
 func deleteDuplicates(head *ListNode) *ListNode {
	 if head==nil{
		 return head
	 }
	 opNode := head
	 for opNode.Next!=nil{
		 if opNode.Val == opNode.Next.Val{
			 opNode.Next = opNode.Next.Next
		 }
	 }
	 return head
}