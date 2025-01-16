package main

import (
	"fmt"
	timers "leetcodes/cmd/leetcodes/timers"
	"os"
	"strconv"
)

type Params struct {
	method string
	size   int
}

const (
	DEFUALT_SIZE   = 1_000_000
	DEFAULT_METHOD = "ALL"
	ITERATIONS     = 1
)

func main() {

	// arg[1] is the method name
	// arg[2] is the input size they want to test

	arguments := parseArgs()
	if arguments.method == "ALL" {
		fmt.Printf("\nTiming All Leetcodes\n")
		timers.TimeAllFunctions(arguments.size, ITERATIONS)
	} else {
		switch arguments.method {
		case "ContainsDuplicate":
			timers.TimeContainsDuplicate(arguments.size, ITERATIONS)
		case "IsAnagram":
			timers.TimeIsAnagram(arguments.size, ITERATIONS)
		case "TwoSum":
			timers.TimeTwoSum(arguments.size, ITERATIONS)
		default:
			fmt.Println("Error: Invalid method name")
			os.Exit(1)
		}
	}
}

func parseArgs() Params {
	var params Params

	if len(os.Args) == 1 {
		params.method = DEFAULT_METHOD
		params.size = DEFUALT_SIZE
	} else if len(os.Args) == 2 {
		params.method = os.Args[1]
		params.size = DEFUALT_SIZE
	} else if len(os.Args) == 3 {
		params.method = os.Args[1]
		m, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Input size must be an integer")
			os.Exit(1)
		} else {
			params.size = m
		}
	} else {
		fmt.Println("Error: Invalid number of arguments")
		os.Exit(1)
	}
	return params
}
