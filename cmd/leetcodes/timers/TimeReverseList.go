package timers

import (
	"fmt"
	leetcodes "leetcodes/LinkedList/ReverseLinkedList"
	"leetcodes/helpers"
	"math/rand"
)

func GetInputReverseList(n int) *leetcodes.ListNode {
	head := &leetcodes.ListNode{}
	dummy := &leetcodes.ListNode{}
	dummy.Next = head
	for i := 0; i < n; i++ {
		val := rand.Intn(n)
		newNode := &leetcodes.ListNode{Val: val}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func TimeReverseList(size int) {
	helpers.PrintHeader("ReverseList", size)

	input := GetInputReverseList(size)

	time_elapsed := helpers.TimeFunctionInputListNodeOutputListNode(leetcodes.ReverseList, input)
	fmt.Printf("\tTime:       %v\n", time_elapsed)

}
