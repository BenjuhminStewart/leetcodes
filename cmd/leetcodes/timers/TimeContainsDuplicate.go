package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/ContainsDuplicate"
	"leetcodes/helpers"
	"math/rand"
)

func getRandomInputContainsDuplicate(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(n * 10)
	}
	return nums
}

func getWorstCaseInputContainsDuplicate(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	return nums
}

func TimeContainsDuplicate(size int) {
	helpers.PrintHeader("ContainsDuplicate", size)

	input := getRandomInputContainsDuplicate(size)
	input_worst := getWorstCaseInputContainsDuplicate(size)

	time_norm := helpers.TimeFunctionInputArrayOutputBool(leetcodes.ContainsDuplicate, input)
	fmt.Printf("\tRandom:     %v\n", time_norm)

	time_worst := helpers.TimeFunctionInputArrayOutputBool(leetcodes.ContainsDuplicate, input_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
