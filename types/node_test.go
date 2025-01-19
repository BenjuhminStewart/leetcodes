package types

import (
	"testing"
)

type createListNodeTestCase struct {
	input    []int
	expected *ListNode
}

type generateSortedListTestCase struct {
	input    int
	expected int
}

func TestCreateListNode(t *testing.T) {
	testCases := []createListNodeTestCase{
		{nil, nil},
		{[]int{1, 2, 3, 4, 5}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}},
	}

	for _, tc := range testCases {
		got := CreateListNode(tc.input...)
		if !got.IsEqualTo(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected.Stringify(), got.Stringify())
		}
	}
}

func TestGenerateSortedList(t *testing.T) {
	t.Parallel()
	testCases := []generateSortedListTestCase{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, 10},
	}

	for _, tc := range testCases {
		got := GenerateSortedList(tc.input)
		if got.Length() != tc.expected {
			t.Errorf("Expected length of %v, got %v", tc.expected, got.Length())
		}
		if !isListNodeSorted(got) {
			t.Errorf("Expected list to be sorted, got %v", got.Stringify())
		}
	}
}

func TestGenerateRandomList(t *testing.T) {
	testCases := []generateSortedListTestCase{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, 10},
	}
	for _, tc := range testCases {
		got := GenerateRandomList(tc.input)
		if got.Length() != tc.expected {
			t.Errorf("Expected length of %v, got %v", tc.expected, got.Length())
		}
	}
}

type stringifyListNodeTestCase struct {
	input    *ListNode
	expected string
}

func TestListNodeStringify(t *testing.T) {
	testCases := []stringifyListNodeTestCase{
		{nil, "nil"},
		{CreateListNode(1), "ListNode: 1 "},
		{CreateListNode(1, 2, 3, 4, 5), "ListNode: 1 2 3 4 5 "},
	}

	for _, tc := range testCases {
		got := tc.input.Stringify()
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

type isEqualToListNodeTestCase struct {
	input1   *ListNode
	input2   *ListNode
	expected bool
}

func TestListNodeIsEqualTo(t *testing.T) {
	t.Parallel()
	testCases := []isEqualToListNodeTestCase{
		{nil, nil, true},
		{CreateListNode(1), CreateListNode(1), true},
		{CreateListNode(1, 2, 3, 4, 5), CreateListNode(1, 2, 3, 4, 5), true},
		{CreateListNode(1, 2, 3, 4, 5), CreateListNode(1, 2, 3, 5, 6), false},
		{CreateListNode(1, 2, 3, 4, 5), CreateListNode(1, 2, 3, 4, 5, 6), false},
	}

	for _, tc := range testCases {
		got := tc.input1.IsEqualTo(tc.input2)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func isListNodeSorted(head *ListNode) bool {
	if head == nil {
		return true
	}
	for head.Next != nil {
		if head.Val > head.Next.Val {
			return false
		}
		head = head.Next
	}
	return true
}
