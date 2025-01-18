package leetcodes

import (
	"testing"
)

type minCostClimbingStairsTestCase struct {
	input    []int
	expected int
}

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []minCostClimbingStairsTestCase{
		{[]int{}, 0},
		{[]int{10}, 0},
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tc := range testCases {
		got := MinCostClimbingStairs(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
