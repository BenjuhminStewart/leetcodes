package leetcodes

import (
	"fmt"
	"testing"
)

type reverseLinkedListTestCase struct {
	input    *ListNode
	expected *ListNode
}

func TestReverseLinkedList(t *testing.T) {
	testCases := []reverseLinkedListTestCase{
		{nil, nil},
		{createListNode(1), createListNode(1)},
		{createListNode(1, 2, 3, 4, 5), createListNode(5, 4, 3, 2, 1)},
		{createListNode(1, 3, 5, 4, 2, 6, 7, 8), createListNode(8, 7, 6, 2, 4, 5, 3, 1)},
	}
	for _, tc := range testCases {
		got := ReverseList(tc.input)

		fmt.Printf("Input: \n")
		printListNode(got)
		fmt.Printf("\nExpected: \n")
		printListNode(tc.expected)
		fmt.Printf("\n")
		if !areListNodesEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func createListNode(vals ...int) *ListNode {
	head := &ListNode{}
	dummy := &ListNode{}
	dummy.Next = head
	for _, val := range vals {
		newNode := &ListNode{Val: val}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func areListNodesEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == b
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}
