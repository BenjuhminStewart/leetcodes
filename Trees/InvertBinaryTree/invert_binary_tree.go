package leetcodes

import "leetcodes/types"

func InvertTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
