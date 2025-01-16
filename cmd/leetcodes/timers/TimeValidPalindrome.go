package timers

import (
	"fmt"
	leetcodes "leetcodes/TwoPointers/ValidPalindrome"
	"leetcodes/helpers"
	"math/rand"
	"strings"
)

func GetRandomInputValidPalindrome(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func GetWorstCaseInputValidPalindrome(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune('a')
	}
	return sb.String()
}

func TimeValidPalindrome(size int) {
	helpers.PrintHeader("ValidPalindrome", size)

	input := GetRandomInputValidPalindrome(size)
	input_worst := GetWorstCaseInputValidPalindrome(size)

	time_elapsed := helpers.TimeFunctionInputStringOutputBool(leetcodes.IsPalindrome, input)
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputStringOutputBool(leetcodes.IsPalindrome, input_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
