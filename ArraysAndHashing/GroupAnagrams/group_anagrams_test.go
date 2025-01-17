package leetcodes

import (
	"sort"
	"testing"
)

type groupAnagramsTestCase struct {
	input    []string
	expected [][]string
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []groupAnagramsTestCase{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"eat", "tea", "ate"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
		{[]string{"a", "b"}, [][]string{{"a"}, {"b"}}},
	}

	for _, tc := range testCases {
		got := GroupAnagrams(tc.input)
		if !are2DArraysEqual(got, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func areArrayElementsEqual(a, b []string) bool {
	set := make(map[string]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		_, containsVal := set[v]
		// remove the value from the set
		if containsVal {
			delete(set, v)
		}
		if !containsVal {
			return false
		}
	}

	// return true if the set is empty
	return len(set) == 0
}

func are2DArraysEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})

	for i := range a {
		if !areArrayElementsEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}
