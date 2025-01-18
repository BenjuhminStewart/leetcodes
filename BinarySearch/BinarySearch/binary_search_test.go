package leetcodes

import (
	"testing"
)

type binarySearchTestCase struct {
	input    []int
	target   int
	expected int
}

func TestBinarySearch(t *testing.T) {
	testCases := []binarySearchTestCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
		{[]int{1, 2, 4, 4, 5, 6, 7, 9, 9, 10}, 3, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 7, 8, 10}, 10, 7},
		{[]int{}, 1, -1},
	}
	for _, tc := range testCases {
		got := BinarySearch(tc.input, tc.target)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
