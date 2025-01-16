package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/ContainsDuplicate"
	"leetcodes/helpers"
	"math/rand"
	"time"
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

func TimeContainsDuplicate(size, iterations int) {
	fmt.Printf("\nTiming 'ContainsDuplicate' | n => %v\n\n", size)
	input := getRandomInputContainsDuplicate(size)
	input_worst := getWorstCaseInputContainsDuplicate(size)
	times := make([]time.Duration, iterations)
	times_worst := make([]time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.ContainsDuplicate(input)
		times[i] = time.Since(<-timer1.C)
		timer1.Stop()
		timer2 := time.NewTimer(time.Second)
		leetcodes.ContainsDuplicate(input_worst)
		times_worst[i] = time.Since(<-timer2.C)
		timer2.Stop()
	}
	fmt.Printf("\tRandom:     %v\n", helpers.GetAverageTime(times))
	fmt.Printf("\tWorst Case: %v\n", helpers.GetAverageTime(times_worst))
}
