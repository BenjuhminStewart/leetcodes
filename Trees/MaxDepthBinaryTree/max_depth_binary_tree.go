package leetcodes

import "leetcodes/types"

func MaxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := MaxDepth(root.Left)
	maxRight := MaxDepth(root.Right)
	if maxLeft > maxRight {
		return maxLeft + 1
	} else {
		return maxRight + 1
	}
}
