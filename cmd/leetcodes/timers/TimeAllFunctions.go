package timers

import (
	"fmt"

	"github.com/fatih/color"
)

func TimeAllFunctions(n int) {
	TimeArrayAndHashingFunctions(n)
	TimeTwoPointersFunctions(n)
	TimeSlidingWindowFunctions(n)
	TimeBinarySearchFunctions(n)
}

func TimeArrayAndHashingFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nArray and Hashing Functions\n")
	fmt.Printf("%v", title)
	TimeContainsDuplicate(n)
	TimeIsAnagram(n)
	TimeTwoSum(n)
	TimeGroupAnagrams(n)
}

func TimeTwoPointersFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nTwo Pointers Functions\n")
	fmt.Printf("%v", title)
	TimeValidPalindrome(n)
}

func TimeSlidingWindowFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nSliding Window Functions\n")
	fmt.Printf("%v", title)
	TimeBuyAndSellStocks(n)
}

func TimeBinarySearchFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nBinary Search Functions\n")
	fmt.Printf("%v", title)
	TimeBinarySearch(n)
}
