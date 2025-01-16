package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/IsAnagram"
	"leetcodes/helpers"
	"math/rand"
	"time"
)

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

func TimeIsAnagram(size, iterations int) {
	fmt.Printf("\nTiming 'IsAnagram' | n => %v\n\n", size)
	input := getRandomInputIsAnagram(size)
	input_worst := getWorstCaseInputIsAnagram(size)

	times := make([]time.Duration, iterations)
	times_worst := make([]time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		timer1 := time.NewTimer(time.Second)
		leetcodes.IsAnagram(input[0], input[1])
		times[i] = time.Since(<-timer1.C)
		timer1.Stop()
		timer2 := time.NewTimer(time.Second)
		leetcodes.IsAnagram(input_worst[0], input_worst[1])
		times_worst[i] = time.Since(<-timer2.C)
		timer2.Stop()
	}
	fmt.Printf("\tRandom:     %v\n", helpers.GetAverageTime(times))
	fmt.Printf("\tWorst Case: %v\n", helpers.GetAverageTime(times_worst))
}
