package timers

import (
	"fmt"
	leetcodes "leetcodes/1DDP/HouseRobber"
	"leetcodes/helpers"
	"math/rand"
)

func GetInputHouseRobber(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(400)
	}
	return arr
}

func TimeHouseRobber(size int) {
	helpers.PrintHeader("HouseRobber", size)

	input := GetInputHouseRobber(size)

	time_elapsed := helpers.TimeFunctionInputIntArrayOutputInt(leetcodes.HouseRobber, input)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
