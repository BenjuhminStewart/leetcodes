package leetcodes

import (
	"leetcodes/types"
	"testing"
)

type maxDepthBinaryTreeTestCase struct {
	input    *types.TreeNode
	expected int
}

func TestMaxDepth(t *testing.T) {
	t.Parallel()
	tests := []maxDepthBinaryTreeTestCase{
		{nil, 0},
		{types.CreateTreeNode(1, nil, nil), 1},
		{types.CreateTreeNode(1, types.CreateTreeNode(2, nil, nil), nil), 2},
		{types.CreateTreeNode(1, types.CreateTreeNode(2, nil, nil), types.CreateTreeNode(3, nil, nil)), 2},
		{types.CreateTreeNode(1, types.CreateTreeNode(2, nil, nil), types.CreateTreeNode(3, types.CreateTreeNode(4, nil, nil), nil)), 3},
		{types.CreateTreeNode(1, types.CreateTreeNode(2, nil, nil), types.CreateTreeNode(3, types.CreateTreeNode(4, nil, nil), types.CreateTreeNode(5, nil, nil))), 3},
	}
	for _, tc := range tests {
		got := MaxDepth(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
