package leetcodes

import (
	"testing"
)

type topKFrequentElementsTestCase struct {
	input_arr []int
	k         int
	expected  []int
}

func TestTopKFrequentElements(t *testing.T) {
	testCases := []topKFrequentElementsTestCase{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 6, 7, 8, 8}, 5, []int{1, 2, 4, 5, 8}},
	}
	for _, tc := range testCases {
		got := TopKFrequentElements(tc.input_arr, tc.k)
		if !areArraysEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func areArraysEqual(a, b []int) bool {
	set := make(map[int]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	for _, v := range b {
		_, containsVal := set[v]
		if !containsVal {
			return false
		}
		delete(set, v)
	}

	return len(set) == 0
}
