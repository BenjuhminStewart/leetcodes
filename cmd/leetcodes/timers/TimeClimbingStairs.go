package timers

import (
	"fmt"
	leetcodes "leetcodes/1DDP/ClimbingStairs"
	"leetcodes/helpers"

	"github.com/fatih/color"
)

const (
	MAX_N_CLIMBING_STAIRS = 91
)

func GetRandomInputClimbingStairs(n int) int {
	return n
}

func GetWorstCaseInputClimbingStairs(n int) int {
	return n
}

func TimeClimbingStairs(size int) {
	actual_size := size
	if size > MAX_N_CLIMBING_STAIRS {
		actual_size = MAX_N_CLIMBING_STAIRS
	}
	helpers.PrintHeader("ClimbingStairs", actual_size)
	if size > MAX_N_CLIMBING_STAIRS {
		color.Red("*N is limited to %v due to surpassing 64 bit integer size*\n\n", MAX_N_CLIMBING_STAIRS)
	}

	input_worst := GetWorstCaseInputClimbingStairs(actual_size)

	time_worst := helpers.TimeFunctionInputIntOutputInt(leetcodes.ClimbingStairs, input_worst)
	fmt.Printf("\tTime:       %v\n", time_worst)
}
