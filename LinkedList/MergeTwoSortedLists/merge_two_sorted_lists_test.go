package leetcode

import (
	"fmt"
	"testing"
)

type mergeTwoSortedListsTestCase struct {
	input1   *ListNode
	input2   *ListNode
	expected *ListNode
}

func TestMergeTwoSortedLists(t *testing.T) {
	t.Parallel()
	testCases := []mergeTwoSortedListsTestCase{
		{nil, nil, nil},
		{createListNode(1), nil, createListNode(1)},
		{nil, createListNode(1), createListNode(1)},
		{createListNode(1), createListNode(2), createListNode(1, 2)},
		{createListNode(1, 3), createListNode(2, 4), createListNode(1, 2, 3, 4)},
		{createListNode(1, 4, 5), createListNode(1, 3, 4), createListNode(1, 1, 3, 4, 4, 5)},
		{createListNode(2, 6, 8), createListNode(4, 7, 9), createListNode(2, 4, 6, 7, 8, 9)},
	}

	for _, tc := range testCases {
		got := MergeTwoLists(tc.input1, tc.input2)
		if !areListNodesEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", StringifyListNode(tc.expected), StringifyListNode(got))
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

func StringifyListNode(head *ListNode) string {
	s := ""
	if head == nil {
		return "nil"
	}
	for head != nil {
		s += fmt.Sprintf("%v ", head.Val)
		head = head.Next
	}
	return s
}
