package timers

import (
	"fmt"
	leetcodes "leetcodes/1D-DP/MinCostClimbingStairs"
	"leetcodes/helpers"
	"math/rand"
)

func GetRandomInputMinCostClimbingStairs(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(1000)
	}
	return res
}

func TimeMinCostClimbingStairs(size int) {
	helpers.PrintHeader("MinCostClimbingStairs", size)

	input := GetRandomInputMinCostClimbingStairs(size)

	time_elapased := helpers.TimeFunctionInputIntArrayOutputInt(leetcodes.MinCostClimbingStairs, input)
	fmt.Printf("\tTime:       %v\n", time_elapased)

}
