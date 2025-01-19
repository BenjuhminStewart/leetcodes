package leetcode

import (
	"leetcodes/types"
	"testing"
)

type mergeTwoSortedListsTestCase struct {
	input1   *types.ListNode
	input2   *types.ListNode
	expected *types.ListNode
}

func TestMergeTwoSortedLists(t *testing.T) {
	t.Parallel()
	testCases := []mergeTwoSortedListsTestCase{
		{nil, nil, nil},
		{types.CreateListNode(1), nil, types.CreateListNode(1)},
		{nil, types.CreateListNode(1), types.CreateListNode(1)},
		{types.CreateListNode(1), types.CreateListNode(2), types.CreateListNode(1, 2)},
		{types.CreateListNode(1, 3), types.CreateListNode(2, 4), types.CreateListNode(1, 2, 3, 4)},
		{types.CreateListNode(1, 4, 5), types.CreateListNode(1, 3, 4), types.CreateListNode(1, 1, 3, 4, 4, 5)},
		{types.CreateListNode(2, 6, 8), types.CreateListNode(4, 7, 9), types.CreateListNode(2, 4, 6, 7, 8, 9)},
	}

	for _, tc := range testCases {
		got := MergeTwoLists(tc.input1, tc.input2)
		if !got.IsEqualTo(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected.Stringify(), got.Stringify())
		}
	}

}
