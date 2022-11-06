package DesignALinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	fastPtr := head
	slowPtr := head

	for fastPtr != nil && fastPtr.Next != nil {
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	return false
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		nextPtr := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextPtr
	}
	return prev
}
