package leetcode

import leetcodes "leetcodes/LinkedList/ReverseLinkedList"

func MergeTwoLists(l1 *leetcodes.ListNode, l2 *leetcodes.ListNode) *leetcodes.ListNode {
	mergedList := &leetcodes.ListNode{}
	node := mergedList

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}

	return mergedList.Next
}
