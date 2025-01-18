package timers

import (
	"fmt"
	leetcodes "leetcodes/1D-DP/ClimbingStairs"
	"leetcodes/helpers"

	"github.com/fatih/color"
)

const (
	MAX_N = 91
)

func GetRandomInputClimbingStairs(n int) int {
	return n
}

func GetWorstCaseInputClimbingStairs(n int) int {
	return n
}

func TimeClimbingStairs(size int) {
	var actual_size int
	if size > MAX_N {
		actual_size = MAX_N
	}
	helpers.PrintHeader("ClimbingStairs", actual_size)
	if size > MAX_N {
		color.Red("N is limited to %v for this function\n\n", MAX_N)
	}

	input_worst := GetWorstCaseInputClimbingStairs(actual_size)

	time_worst := helpers.TimeFunctionInputIntOutputInt(leetcodes.ClimbingStairs, input_worst)
	fmt.Printf("\tTime: %v\n", time_worst)
}
