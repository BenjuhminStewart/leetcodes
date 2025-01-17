package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/GroupAnagrams"
	"leetcodes/helpers"
	"math/rand"
)

func GetRandomInputGroupAnagrams(n int) []string {
	res := make([]string, n)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := range res {
		word_length := rand.Intn(5) + 1
		curr := make([]byte, word_length)
		for j := 0; j < word_length; j++ {
			curr[j] = letters[rand.Intn(26)]
		}
		res[i] = string(curr)
	}

	return res
}

func GetWorstCaseInputGroupAnagrams(n int) []string {
	res := make([]string, n)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := range res {
		word_length := 5
		curr := make([]byte, word_length)
		for j := 0; j < word_length; j++ {
			curr[j] = letters[rand.Intn(26)]
		}
		res[i] = string(curr)
	}
	return res
}

func TimeGroupAnagrams(size int) {
	helpers.PrintHeader("GroupAnagrams", size)

	input := GetRandomInputGroupAnagrams(size)
	input_worst := GetWorstCaseInputGroupAnagrams(size)

	time_elapsed := helpers.TimeFunctionInputStringArrayOutputStringArrayArray(leetcodes.GroupAnagrams, input)
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputStringArrayOutputStringArrayArray(leetcodes.GroupAnagrams, input_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
