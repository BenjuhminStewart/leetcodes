package leetcodes

import (
	"testing"
)

type twoSumTestCase struct {
	input    []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	t.Parallel()
	testCases := []twoSumTestCase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{4, 6, 6, 8}, 20, []int{-1, -1}},
		{[]int{}, 14, []int{-1, -1}},
	}
	for _, tc := range testCases {
		got := TwoSum(tc.input, tc.target)
		if !areArraysEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func areArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
