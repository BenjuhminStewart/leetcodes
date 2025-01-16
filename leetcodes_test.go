package leetcodes_test

import (
	"leetcodes"
	"testing"
)

type containsDuplicateTestCase struct {
	input    []int
	expected bool
}

type isAnagramTestCase struct {
	s        string
	t        string
	expected bool
}

type twoSumTestCase struct {
	input    []int
	target   int
	expected []int
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
		got := leetcodes.ContainsDuplicate(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []isAnagramTestCase{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"hello", "llloh", false},
		{"a", "a", true},
		{"", "", true},
		{"!#@$%&^*)(", "()*^&$%#!@", true},
	}

	for _, tc := range testCases {
		got := leetcodes.IsAnagram(tc.s, tc.t)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
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
		got := leetcodes.TwoSum(tc.input, tc.target)
		if !AreArraysEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func AreArraysEqual(a, b []int) bool {
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
