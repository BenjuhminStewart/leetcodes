package timers

import (
	"fmt"
	leetcodes "leetcodes/LinkedList/MergeTwoSortedLists"
	"leetcodes/helpers"
	"leetcodes/types"
)

func GetInputMergeTwoSortedLists(n int) (*types.ListNode, *types.ListNode) {
	// generate two sorted linked lists
	l1 := types.GenerateSortedList(n)
	l2 := types.GenerateSortedList(n)
	return l1, l2

}

func TimeMergeTwoSortedLists(size int) {
	helpers.PrintHeader("MergeTwoSortedLists", size)
	input1, input2 := GetInputMergeTwoSortedLists(size)
	time_elapsed := helpers.TimeFunctionInputListNodeListNodeOutputListNode(leetcodes.MergeTwoLists, input1, input2)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
