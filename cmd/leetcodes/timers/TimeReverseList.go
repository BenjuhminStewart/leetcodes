package timers

import (
	"fmt"
	leetcodes "leetcodes/LinkedList/ReverseLinkedList"
	"leetcodes/helpers"
	"leetcodes/types"
)

func TimeReverseList(size int) {
	helpers.PrintHeader("ReverseList", size)
	input := types.GenerateRandomList(size)
	time_elapsed := helpers.TimeFunctionInputListNodeOutputListNode(leetcodes.ReverseList, input)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
