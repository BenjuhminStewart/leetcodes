package types

import (
	"testing"
)

type createTreeNodeTestCase struct {
	value    int
	left     *TreeNode
	right    *TreeNode
	expected *TreeNode
}

func TestCreateTreeNode(t *testing.T) {
	testCases := []createTreeNodeTestCase{
		{1, nil, nil, &TreeNode{Val: 1, Left: nil, Right: nil}},
		{2, &TreeNode{Val: 1, Left: nil, Right: nil}, &TreeNode{Val: 3, Left: nil, Right: nil}, &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
	}

	for _, tc := range testCases {
		got := CreateTreeNode(tc.value, tc.left, tc.right)
		if !got.IsEqualTo(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected.Stringify(), got.Stringify())
		}
	}
}

type generateRandomTreeTestCase struct {
	input    int
	expected int
}

func TestGenerateRandomTree(t *testing.T) {
	t.Parallel()
	testCases := []generateRandomTreeTestCase{
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
		got := GenerateRandomTree(tc.input)
		if got.Length() != tc.expected {
			t.Errorf("Expected length of %v, got %v", tc.expected, got.Length())
		}
	}
}

type stringifyTreeNodeTestCase struct {
	input    *TreeNode
	expected string
}

func TestTreeNodeStringify(t *testing.T) {
	testCases := []stringifyTreeNodeTestCase{
		{nil, ""},
		{CreateTreeNode(1, nil, nil), "1 "},
		{CreateTreeNode(1, CreateTreeNode(2, nil, nil), CreateTreeNode(3, nil, nil)), "1 2 3 "},
	}

	for _, tc := range testCases {
		got := tc.input.Stringify()
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

type lengthTreeNodeTestCase struct {
	input    *TreeNode
	expected int
}

func TestTreeNodeLength(t *testing.T) {
	testCases := []lengthTreeNodeTestCase{
		{nil, 0},
		{CreateTreeNode(1, nil, nil), 1},
		{CreateTreeNode(1, CreateTreeNode(2, nil, nil), CreateTreeNode(3, nil, nil)), 3},
		{CreateTreeNode(1, CreateTreeNode(2, CreateTreeNode(4, nil, nil), CreateTreeNode(5, nil, nil)), CreateTreeNode(3, nil, nil)), 5},
	}

	for _, tc := range testCases {
		got := tc.input.Length()
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
