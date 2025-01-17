package leetcodes

import (
	"testing"
)

type buyAndSellStockTestCase struct {
	input    []int
	expected int
}

func TestBuyAndSellStockMaxProfit(t *testing.T) {
	testCases := []buyAndSellStockTestCase{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1, 2, 7, 1, 5, 3, 6, 4}, 6},
		{[]int{100}, 0},
	}
	for _, tc := range testCases {
		got := BuyAndSellStockMaxProfit(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}
