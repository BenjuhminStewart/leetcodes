package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"leetcodes/types"
	"time"
)

func PrintHeader(method string, size int) {
	title := color.New(color.FgYellow, color.Bold).Sprintf("%v", method)
	fmt.Printf("\nTiming %v | n => %v\n\n", title, size)
}

type fn_int_int func(int) int
type fn_arr_bool func([]int) bool
type fn_str_str_bool func(string, string) bool
type fn_intarr_int_intarr func([]int, int) []int
type fn_intarr_int_int func([]int, int) int
type fn_str_bool func(string) bool
type fn_intarr_int func([]int) int
type fn_strarr_strarrarr func([]string) [][]string
type fn_ll_ll func(*types.ListNode) *types.ListNode
type fn_ll_ll_ll func(*types.ListNode, *types.ListNode) *types.ListNode
type fn_tree_tree func(*types.TreeNode) *types.TreeNode
type fn_tree_int func(*types.TreeNode) int

func TimeFunctionInputIntOutputInt(f fn_int_int, input int) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

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

func TimeFunctionInputIntArrayOutputInt(f fn_intarr_int, input []int) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputIntArrayIntOutputInt(f fn_intarr_int_int, input []int, target int) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input, target)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputStringArrayOutputStringArrayArray(f fn_strarr_strarrarr, input []string) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputListNodeOutputListNode(f fn_ll_ll, input *types.ListNode) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputListNodeListNodeOutputListNode(f fn_ll_ll_ll, input1 *types.ListNode, input2 *types.ListNode) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input1, input2)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputTreeNodeOutputTreeNode(f fn_tree_tree, input *types.TreeNode) time.Duration {
	timer := time.NewTimer(time.Second)
	f(input)
	timeElapsed := time.Since(<-timer.C)
	timer.Stop()
	return timeElapsed
}

func TimeFunctionInputTreeNodeOutputInt(f fn_tree_int, input *types.TreeNode) time.Duration {
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
