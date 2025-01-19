package leetcodes

import (
	"testing"
)

type validParenthesesTestCase struct {
	input    string
	expected bool
}

func TestValidParentheses(t *testing.T) {
	testCases := []validParenthesesTestCase{
		{"", true},
		{"ab", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"([)]}", false},
		{"{[()([])]}", true},
		{"{[]}", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"([)]}", false},
		{"{[]}", true},
	}
	for _, tc := range testCases {
		got := IsValid(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
