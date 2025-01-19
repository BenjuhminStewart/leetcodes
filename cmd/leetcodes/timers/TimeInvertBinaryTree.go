package timers

import (
	"fmt"
	leetcodes "leetcodes/Trees/InvertBinaryTree"
	"leetcodes/helpers"
	"leetcodes/types"

	"github.com/fatih/color"
)

const (
	MAX_N_INVERT_BINARY_TREE = 10_000_000
)

func TimeInvertBinaryTree(size int) {
	helpers.PrintHeader("InvertBinaryTree", size)
	if size > MAX_N_INVERT_BINARY_TREE {
		color.Red("*N is limited to %v for this function due to recursion limits*\n\n", MAX_N_INVERT_BINARY_TREE)
		size = MAX_N_INVERT_BINARY_TREE
	}
	input := types.GenerateRandomTree(size)
	time_elapsed := helpers.TimeFunctionInputTreeNodeOutputTreeNode(leetcodes.InvertTree, input)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
