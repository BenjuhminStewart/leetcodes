package timers

import (
	"fmt"
	leetcodes "leetcodes/LinkedList/MergeTwoSortedLists"
	leetcodesNode "leetcodes/LinkedList/ReverseLinkedList"
	"leetcodes/helpers"
	"math/rand"
)

func GetInputMergeTwoSortedLists(n int) (*leetcodesNode.ListNode, *leetcodesNode.ListNode) {
	// generate two sorted linked lists
	l1 := createListNode(n)
	l2 := createListNode(n)
	return l1, l2

}

func createListNode(n int) *leetcodesNode.ListNode {
	head := &leetcodesNode.ListNode{}
	dummy := &leetcodesNode.ListNode{}
	dummy.Next = head
	curr := 0
	for i := 0; i < n; i++ {
		val := rand.Intn(2)
		if val == 0 {
			curr++
		} else {
			curr += 2
		}
		newNode := &leetcodesNode.ListNode{Val: curr}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func TimeMergeTwoSortedLists(size int) {
	helpers.PrintHeader("MergeTwoSortedLists", size)

	input1, input2 := GetInputMergeTwoSortedLists(size)

	time_elapsed := helpers.TimeFunctionInputListNodeListNodeOutputListNode(leetcodes.MergeTwoLists, input1, input2)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}

func printListNode(head *leetcodesNode.ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}
