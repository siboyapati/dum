package MiddleOfLinkedList_876

func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}

type ListNode struct {
	Val int
	Next *ListNode
}