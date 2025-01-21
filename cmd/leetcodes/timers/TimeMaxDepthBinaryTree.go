package timers

import (
	"fmt"
	leetcodes "leetcodes/Trees/MaxDepthBinaryTree"
	"leetcodes/helpers"
	"leetcodes/types"

	"github.com/fatih/color"
)

const (
	MAX_N_MAX_DEPTH_BINARY_TREE = 10_000_000
)

func GetInputMaxDepthBinaryTree(n int) *types.TreeNode {
	return types.GenerateRandomTree(n)
}

func TimeMaxDepthBinaryTree(size int) {
	actual_size := size
	if size >= MAX_N_MAX_DEPTH_BINARY_TREE {
		actual_size = MAX_N_MAX_DEPTH_BINARY_TREE
	}
	helpers.PrintHeader("MaxDepthBinaryTree", actual_size)
	if size >= MAX_N_MAX_DEPTH_BINARY_TREE {
		color.Red("N is limited to %v due to recursion on stack\n\n", MAX_N_MAX_DEPTH_BINARY_TREE)
	}
	input := GetInputMaxDepthBinaryTree(actual_size)
	time_elapsed := helpers.TimeFunctionInputTreeNodeOutputInt(leetcodes.MaxDepth, input)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
