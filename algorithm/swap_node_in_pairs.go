package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy = &ListNode{Val: -1, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		var first = cur.Next
		var second = cur.Next.Next

		first.Next = second.Next
		cur.Next = second
		cur.Next.Next = first

		cur = cur.Next.Next
	}

	return dummy.Next
}
