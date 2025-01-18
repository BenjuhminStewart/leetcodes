package timers

import (
	"fmt"
	leetcodes "leetcodes/BinarySearch/BinarySearch"
	"leetcodes/helpers"
	"math/rand"
)

func GetRandomInputBinarySearch(n int) ([]int, int) {
	nums := make([]int, n)
	last_num := 0
	for i := range nums {
		if rand.Intn(2) == 0 {
			nums[i] = last_num + 1
			last_num = nums[i]
		} else {
			nums[i] = last_num + 2
			last_num = nums[i]
		}
	}
	target := rand.Intn(last_num)
	return nums, target
}

func GetWorstCaseInputBinarySearch(n int) ([]int, int) {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	target := n/2 + 1
	return nums, target
}

func TimeBinarySearch(size int) {
	helpers.PrintHeader("BinarySearch", size)

	input, target := GetRandomInputBinarySearch(size)
	input_worst, target_worst := GetWorstCaseInputBinarySearch(size)

	time_elapsed := helpers.TimeFunctionInputIntArrayIntOutputInt(leetcodes.BinarySearch, input, target)
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputIntArrayIntOutputInt(leetcodes.BinarySearch, input_worst, target_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
