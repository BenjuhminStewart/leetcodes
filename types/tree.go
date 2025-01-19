package types

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func GenerateRandomTree(n int) *TreeNode {
	head := &TreeNode{}
	dummy := &TreeNode{}
	dummy.Left = head
	for i := 0; i < n; i++ {
		rand_val := rand.Intn(2)
		flipper := 1
		if rand_val == 0 {
			flipper = -1
		}
		val := rand.Intn(100) * flipper
		newNode := &TreeNode{Val: val}
		head.Right = newNode
		head = newNode
	}
	return dummy.Left.Right
}

func (node *TreeNode) IsEqualTo(other *TreeNode) bool {
	if node == nil && other == nil {
		return true
	}

	return node.Val == other.Val && node.Left.IsEqualTo(other.Left) && node.Right.IsEqualTo(other.Right)
}

func (node *TreeNode) Stringify() string {
	return stringifyHelper(node, "")
}

func stringifyHelper(node *TreeNode, s string) string {
	if node == nil {
		return s
	}
	s += fmt.Sprintf("%v ", node.Val)
	s = stringifyHelper(node.Left, s)
	s = stringifyHelper(node.Right, s)
	return s
}

func (node *TreeNode) Length() int {
	return lengthHelper(node, 0)
}

func lengthHelper(node *TreeNode, length int) int {
	if node == nil {
		return length
	}
	length += 1
	length = lengthHelper(node.Left, length)
	length = lengthHelper(node.Right, length)
	return length
}
