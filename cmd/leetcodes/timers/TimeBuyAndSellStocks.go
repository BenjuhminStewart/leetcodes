package timers

import (
	"fmt"
	leetcodes "leetcodes/SlidingWindows/BuyAndSellStock"
	"leetcodes/helpers"
	"math/rand"
)

func getRandomInputBuyAndSellStocks(n int) []int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	return prices
}

func getWorstCaseInputBuyAndSellStocks(n int) []int {
	prices := make([]int, n)
	j := 0
	for i := range prices {
		prices[i] = j
		j++
	}
	return prices
}

func TimeBuyAndSellStocks(size int) {
	helpers.PrintHeader("BuyAndSellStocks", size)

	input := getRandomInputBuyAndSellStocks(size)
	input_worst := getWorstCaseInputBuyAndSellStocks(size)

	time_elapsed := helpers.TimeFunctionInputIntArrayOutputInt(leetcodes.BuyAndSellStockMaxProfit, input)
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputIntArrayOutputInt(leetcodes.BuyAndSellStockMaxProfit, input_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
