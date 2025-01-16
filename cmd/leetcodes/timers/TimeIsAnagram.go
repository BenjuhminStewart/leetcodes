package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/IsAnagram"
	"leetcodes/helpers"
	"math/rand"
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

func TimeIsAnagram(size int) {
	helpers.PrintHeader("IsAnagram", size)
	input := getRandomInputIsAnagram(size)
	input_worst := getWorstCaseInputIsAnagram(size)

	time_elapsed := helpers.TimeFunctionInputTwoStringsOutputBool(leetcodes.IsAnagram, input[0], input[1])
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputTwoStringsOutputBool(leetcodes.IsAnagram, input_worst[0], input_worst[1])
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
