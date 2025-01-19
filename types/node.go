package types

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(vals ...int) *ListNode {
	head := &ListNode{}
	dummy := &ListNode{}
	dummy.Next = head
	for _, val := range vals {
		newNode := &ListNode{Val: val}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func GenerateSortedList(n int) *ListNode {
	head := &ListNode{}
	dummy := &ListNode{}
	dummy.Next = head
	curr := 0
	for i := 0; i < n; i++ {
		val := rand.Intn(2)
		if val == 0 {
			curr++
		} else {
			curr += 2
		}
		newNode := &ListNode{Val: curr}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func GenerateRandomList(n int) *ListNode {
	head := &ListNode{}
	dummy := &ListNode{}
	dummy.Next = head
	for i := 0; i < n; i++ {
		val := rand.Intn(n * 10)
		newNode := &ListNode{Val: val}
		head.Next = newNode
		head = newNode
	}
	return dummy.Next.Next
}

func (node *ListNode) Print() {
	if node == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("ListNode: ")
	for node != nil {
		fmt.Printf("%v ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func (node *ListNode) Stringify() string {
	s := ""
	if node == nil {
		return "nil"
	}
	for node != nil {
		s += fmt.Sprintf("%v ", node.Val)
		node = node.Next
	}
	return s
}

func (node *ListNode) IsEqualTo(other *ListNode) bool {
	for node != nil && other != nil {
		if node.Val != other.Val {
			return false
		}
		node = node.Next
		other = other.Next
	}
	return node == other
}
