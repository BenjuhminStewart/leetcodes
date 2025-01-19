package leetcodes

import (
	"leetcodes/types"
	"testing"
)

type reverseLinkedListTestCase struct {
	input    *types.ListNode
	expected *types.ListNode
}

func TestReverseLinkedList(t *testing.T) {
	testCases := []reverseLinkedListTestCase{
		{nil, nil},
		{types.CreateListNode(1), types.CreateListNode(1)},
		{types.CreateListNode(1, 2, 3, 4, 5), types.CreateListNode(5, 4, 3, 2, 1)},
		{types.CreateListNode(1, 3, 5, 4, 2, 6, 7, 8), types.CreateListNode(8, 7, 6, 2, 4, 5, 3, 1)},
	}
	for _, tc := range testCases {
		got := ReverseList(tc.input)

		if !got.IsEqualTo(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
