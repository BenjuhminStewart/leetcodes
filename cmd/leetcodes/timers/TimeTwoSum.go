package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/TwoSum"
	"leetcodes/helpers"
	"math/rand"
	"time"
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

func TimeTwoSum(n, iterations int) {
	fmt.Printf("\nTiming 'TwoSum' | n => %v\n\n", n)
	arr, target := getRandomInputTwoSum(n)
	arr_worst, target_worst := getWorstCaseInputTwoSum(n)

	times := make([]time.Duration, iterations)
	times_worst := make([]time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.TwoSum(arr, target)
		times[i] = time.Since(<-timer1.C)
		timer1.Stop()
		timer2 := time.NewTimer(time.Second)
		leetcodes.TwoSum(arr_worst, target_worst)
		times_worst[i] = time.Since(<-timer2.C)
		timer2.Stop()
	}
	fmt.Printf("\tRandom:      %v\n", helpers.GetAverageTime(times))
	fmt.Printf("\tWorst Case:  %v\n", helpers.GetAverageTime(times_worst))
}
