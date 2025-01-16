package main

import (
	"fmt"
	"leetcodes"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	ITERATIONS = 1
)

func main() {
	manage() // Launch a CPU-bound task
}

func manage() {
	n := 0
	arglen := len(os.Args)
	if arglen == 1 {
		n = 10_000_000
	} else if arglen == 2 {
		m, err := strconv.Atoi(os.Args[1])
		if err != nil {
			err_message := fmt.Errorf("must input a number")
			fmt.Println(err_message)
			os.Exit(1)
		} else {
			n = m
		}
	} else {
		t, err2 := strconv.Atoi(os.Args[2])
		m, err1 := strconv.Atoi(os.Args[1])
		if err1 != nil {
			err1_message := fmt.Errorf("must input a number")
			fmt.Println(err1_message)
			os.Exit(1)

		} else if err2 != nil {
			err2_message := fmt.Errorf("must input a number")
			fmt.Println(err2_message)
			os.Exit(1)
		} else {
			n = m
			ITERATIONS = t

		}
	}
	fmt.Printf("Size of Input: %v\nIterations: %v\n---------------------------------------------\n", n, ITERATIONS)
	timeAllFunctions(n)
}

func timeAllFunctions(n int) {
	timeContainsDuplicate(n)
	timeIsAnagram(n)
	timeTwoSum(n)
}

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

func getAverage(times []time.Duration) time.Duration {
	sum := 0.0
	for _, v := range times {
		sum += float64(v.Microseconds())
	}
	return time.Duration(sum/float64(len(times))) * 1000
}

func timeContainsDuplicate(size int) {

	input := getRandomInputContainsDuplicate(size)
	input_worst := getWorstCaseInputContainsDuplicate(size)

	// start timer
	times := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.ContainsDuplicate(input)
		times[i] = time.Since(<-timer1.C)
		timer1.Stop()
	}
	fmt.Printf("ContainsDuplicate [Random | %v]: %v\n", size, getAverage(times))

	times_worst := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer2 := time.NewTimer(time.Second)
		leetcodes.ContainsDuplicate(input_worst)
		times_worst[i] = time.Since(<-timer2.C)
		timer2.Stop()
	}

	fmt.Printf("ContainsDuplicate [Worst Case | %v]: %v\n\n", size, getAverage(times_worst))
}

func getRandomInputIsAnagram(n int) []string {
	letters := "abcdefghijklmnopqrstuvwxyz!@#$%^&*()1234567890"
	res := make([]string, 2)
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	res[0] = string(s)
	s = make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	res[1] = string(s)
	return res
}

func getWorstCaseInputIsAnagram(n int) []string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	res := make([]string, 2)
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	res[0] = string(s)
	res[1] = string(s)
	return res
}

func timeIsAnagram(size int) {
	input := getRandomInputIsAnagram(size)
	input_worst := getWorstCaseInputIsAnagram(size)

	times := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.IsAnagram(input[0], input[1])
		times[i] = time.Since(<-timer1.C)
		timer1.Stop()
	}
	fmt.Printf("IsAnagram [Random | %v]: %v\n", size, getAverage(times))

	times_worst := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer2 := time.NewTimer(time.Second)
		leetcodes.IsAnagram(input_worst[0], input_worst[1])
		times_worst[i] = time.Since(<-timer2.C)
		timer2.Stop()
	}
	fmt.Printf("IsAnagram [Worst Case | %v]: %v\n\n", size, getAverage(times_worst))
}

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

func timeTwoSum(n int) {
	arr, target := getRandomInputTwoSum(n)
	arr_worst, target_worst := getWorstCaseInputTwoSum(n)

	times := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.TwoSum(arr, target)
		times[i] = time.Since(<-timer1.C)
	}
	fmt.Printf("TwoSum [Random | %v]: %v\n", n, getAverage(times))

	times_worst := make([]time.Duration, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		timer2 := time.NewTimer(time.Second)
		leetcodes.TwoSum(arr_worst, target_worst)
		times_worst[i] = time.Since(<-timer2.C)
	}
	fmt.Printf("TwoSum [Worst Case | %v]: %v\n", n, getAverage(times_worst))
}
