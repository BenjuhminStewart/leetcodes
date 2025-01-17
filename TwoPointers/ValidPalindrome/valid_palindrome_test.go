package leetcodes

import (
	"testing"
)

type validPalindromeTestCase struct {
	input    string
	expected bool
}

func TestValidPalindrome(t *testing.T) {
	testCases := []validPalindromeTestCase{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"racecar", true},
		{" ", true},
		{"", true},
	}
	for _, tc := range testCases {
		got := IsPalindrome(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
