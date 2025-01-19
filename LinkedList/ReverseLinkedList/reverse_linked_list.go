package leetcodes

import "leetcodes/types"

func ReverseList(head *types.ListNode) *types.ListNode {
	var prev *types.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
