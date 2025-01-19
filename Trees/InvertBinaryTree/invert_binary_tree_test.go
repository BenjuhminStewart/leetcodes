package leetcodes

import (
	"leetcodes/types"
	"testing"
)

type invertBinaryTreeTestCase struct {
	input    *types.TreeNode
	expected *types.TreeNode
}

func TestInvertBinaryTree(t *testing.T) {

	input1 := types.CreateTreeNode(4, types.CreateTreeNode(2, types.CreateTreeNode(1, nil, nil), types.CreateTreeNode(3, nil, nil)), types.CreateTreeNode(7, types.CreateTreeNode(6, nil, nil), types.CreateTreeNode(9, nil, nil)))

	output1 := types.CreateTreeNode(4, types.CreateTreeNode(7, types.CreateTreeNode(9, nil, nil), types.CreateTreeNode(6, nil, nil)), types.CreateTreeNode(2, types.CreateTreeNode(3, nil, nil), types.CreateTreeNode(1, nil, nil)))

	input2 := types.CreateTreeNode(2, types.CreateTreeNode(1, nil, nil), types.CreateTreeNode(3, nil, nil))

	output2 := types.CreateTreeNode(2, types.CreateTreeNode(3, nil, nil), types.CreateTreeNode(1, nil, nil))

	testCases := []invertBinaryTreeTestCase{
		{input1, output1},
		{input2, output2},
	}
	for _, tc := range testCases {
		got := InvertTree(tc.input)
		if !got.IsEqualTo(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected.Stringify(), got.Stringify())
		}
	}
}
