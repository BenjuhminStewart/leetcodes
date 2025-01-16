package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/TwoSum"
	"leetcodes/helpers"
	"math/rand"
)

func getRandomInputTwoSum(n int) ([]int, int) {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(n * 10)

	}
	return nums, nums[rand.Intn(n)] * 10
}

func getWorstCaseInputTwoSum(n int) ([]int, int) {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	return nums, -1
}

func TimeTwoSum(n int) {
	helpers.PrintHeader("TwoSum", n)

	arr, target := getRandomInputTwoSum(n)
	arr_worst, target_worst := getWorstCaseInputTwoSum(n)

	time_elapsed := helpers.TimeFunctionInputIntArrayIntOutputIntArray(leetcodes.TwoSum, arr, target)
	fmt.Printf("\tRandom:      %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputIntArrayIntOutputIntArray(leetcodes.TwoSum, arr_worst, target_worst)
	fmt.Printf("\tWorst Case:  %v\n", time_worst)
}
