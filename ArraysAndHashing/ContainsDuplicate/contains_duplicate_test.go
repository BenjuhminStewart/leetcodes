package leetcodes

import (
	"testing"
)

type containsDuplicateTestCase struct {
	input    []int
	expected bool
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []containsDuplicateTestCase{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, true},
		{[]int{1}, false},
	}

	for _, tc := range testCases {
		got := ContainsDuplicate(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
