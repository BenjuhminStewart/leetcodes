package timers

import (
	"fmt"
	leetcodes "leetcodes/Stack/ValidParentheses"
	"leetcodes/helpers"
	"math/rand"
	"strings"
)

func GetRandomInputValidParentheses(n int) string {
	letters := "(){}[]"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func GetWorstCaseInputValidParentheses(n int) string {
	var sb strings.Builder
	for i := 0; i < n/2; i++ {
		sb.WriteRune('(')
		sb.WriteRune(')')
	}
	return sb.String()
}

func TimeValidParentheses(size int) {
	helpers.PrintHeader("ValidParentheses", size)

	input := GetRandomInputValidParentheses(size)
	input_worst := GetWorstCaseInputValidParentheses(size)

	time_elapsed := helpers.TimeFunctionInputStringOutputBool(leetcodes.IsValid, input)
	fmt.Printf("\tRandom:     %v\n", time_elapsed)

	time_worst := helpers.TimeFunctionInputStringOutputBool(leetcodes.IsValid, input_worst)
	fmt.Printf("\tWorst Case: %v\n", time_worst)
}
