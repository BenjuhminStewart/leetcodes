package leetcodes

import (
	"testing"
)

type houseRobberTestCase struct {
	input    []int
	expected int
}

func TestHouseRobber(t *testing.T) {
	testCases := []houseRobberTestCase{
		{[]int{}, 0},
		{[]int{100}, 100},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, tc := range testCases {
		got := HouseRobber(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
