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
)

func main() {

	// arg[1] is the method name
	// arg[2] is the input size they want to test

	arguments := parseArgs()
	switch arguments.method {
	case "ALL":
		timers.TimeAllFunctions(arguments.size)
	case "ArraysAndHashing":
		timers.TimeArrayAndHashingFunctions(arguments.size)
	case "TwoPointers":
		timers.TimeTwoPointersFunctions(arguments.size)
	case "SlidingWindows":
		timers.TimeSlidingWindowFunctions(arguments.size)
	case "BinarySearchFunctions":
		timers.TimeBinarySearchFunctions(arguments.size)
	case "1DDP":
		timers.Time1DDPFunctions(arguments.size)
	case "LinkedList":
		timers.TimeLinkedListFunctions(arguments.size)
	case "Tree":
		timers.TimeTreeFunctions(arguments.size)
	case "ContainsDuplicate":
		timers.TimeContainsDuplicate(arguments.size)
	case "IsAnagram":
		timers.TimeIsAnagram(arguments.size)
	case "TwoSum":
		timers.TimeTwoSum(arguments.size)
	case "ValidPalindrome":
		timers.TimeValidPalindrome(arguments.size)
	case "BuyAndSellStocks":
		timers.TimeBuyAndSellStocks(arguments.size)
	case "GroupAnagrams":
		timers.TimeGroupAnagrams(arguments.size)
	case "BinarySearch":
		timers.TimeBinarySearch(arguments.size)
	case "ClimbingStairs":
		timers.TimeClimbingStairs(arguments.size)
	case "MinCostClimbingStairs":
		timers.TimeMinCostClimbingStairs(arguments.size)
	case "HouseRobber":
		timers.TimeHouseRobber(arguments.size)
	case "ReverseLinkedList":
		timers.TimeReverseList(arguments.size)
	case "MergeTwoSortedLists":
		timers.TimeMergeTwoSortedLists(arguments.size)
	case "InvertBinaryTree":
		timers.TimeInvertBinaryTree(arguments.size)
	case "help":
		printHelp()
	default:
		fmt.Println("Error: Invalid method name")
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Usage: leetcodes [method] [input_size_n]")
	fmt.Println("Available methods:")
	fmt.Println("- ALL")
	fmt.Println("- ArraysAndHashing")
	fmt.Println("  - ContainsDuplicate")
	fmt.Println("  - IsAnagram")
	fmt.Println("  - TwoSum")
	fmt.Println("- TwoPointers")
	fmt.Println("  - ValidPalindrome")
	fmt.Println("- SlidingWindows")
	fmt.Println("  - BuyAndSellStocks")
	fmt.Println("- BinarySearchFunctions")
	fmt.Println("  - BinarySearch")
	fmt.Println("- 1DDP")
	fmt.Println("  - ClimbingStairs")
	fmt.Println("  - MinCostClimbingStairs")
	fmt.Println("  - HouseRobber")
	fmt.Println("- LinkedList")
	fmt.Println("  - ReverseLinkedList")
	fmt.Println("  - MergeTwoSortedLists")
	fmt.Println("- Tree")
	fmt.Println("  - InvertBinaryTree")
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
