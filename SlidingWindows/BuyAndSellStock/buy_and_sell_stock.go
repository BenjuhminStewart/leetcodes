package leetcodes

func BuyAndSellStockMaxProfit(prices []int) int {
	max_profit := 0
	for low, high := 0, 1; high < len(prices); high++ {
		if prices[high] < prices[low] {
			low = high
		} else {
			max_profit = max(max_profit, prices[high]-prices[low])
		}
	}

	return max_profit
}
