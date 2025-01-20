package timers

import (
	"fmt"
	leetcodes "leetcodes/ArraysAndHashing/TopKFrequentElements"
	"leetcodes/helpers"
	"math/rand"
)

const (
	MAX_VAL_LEETCODE_TOP_K_FREQUENT_ELEMENTS = 10_000
)

func GetInputTopKFrequentElements(n int) ([]int, int) {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(MAX_VAL_LEETCODE_TOP_K_FREQUENT_ELEMENTS)
	}
	return arr, getValidK(arr)
}

func getValidK(arr []int) int {
	set := make(map[int]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return len(set)
}

func TimeTopKFrequentElements(size int) {
	helpers.PrintHeader("TopKFrequentElements", size)
	input, k := GetInputTopKFrequentElements(size)
	time_elapsed := helpers.TimeFunctionInputIntArrayIntOutputIntArray(leetcodes.TopKFrequentElements, input, k)
	fmt.Printf("\tTime:       %v\n", time_elapsed)
}
