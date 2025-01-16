package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func PrintHeader(method string, size int) {
	title := color.New(color.FgYellow, color.Bold).Sprintf("%v", method)
	fmt.Printf("\nTiming %v | n => %v\n\n", title, size)
}

type fn_arr_bool func([]int) bool
type fn_str_str_bool func(string, string) bool
type fn_intarr_int_intarr func([]int, int) []int
type fn_str_bool func(string) bool

func TimeFunctionInputArrayOutputBool(f fn_arr_bool, input []int) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputTwoStringsOutputBool(f fn_str_str_bool, input_1 string, input_2 string) time.Duration {

	timer := time.NewTimer(time.Second)
	f(input_1, input_2)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputIntArrayIntOutputIntArray(f fn_intarr_int_intarr, input []int, target int) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input, target)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputStringOutputBool(f fn_str_bool, input string) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func GetAverageTime(times []time.Duration) time.Duration {
	sum := 0.0
	for _, v := range times {
		sum += float64(v.Microseconds())
	}
	return time.Duration(sum/float64(len(times))) * 1000
}
