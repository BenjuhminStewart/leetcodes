package leetcodes

import (
	"testing"
)

type isAnagramTestCase struct {
	s        string
	t        string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	testCases := []isAnagramTestCase{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"hello", "llloh", false},
		{"a", "a", true},
		{"", "", true},
		{"!#@$%&^*)(", "()*^&$%#!@", true},
		{"ab", "a", false},
	}

	for _, tc := range testCases {
		got := IsAnagram(tc.s, tc.t)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
