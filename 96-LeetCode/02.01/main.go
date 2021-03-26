package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// fmt.Println(removeDuplicateNodes())
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	var out *ListNode
	mapHas := make(map[int]bool)
	if head != nil {
		out = new(ListNode)
		out.Val = head.Val
		mapHas[head.Val] = true
	} else {
		return out
	}
	tem := out
	for head != nil {
		_, ok := mapHas[head.Val]
		if !ok {
			tem.Next = new(ListNode)
			tem.Next.Val = head.Val
			tem = tem.Next
			mapHas[head.Val] = true
		}
		head = head.Next
	}

	return out
}
