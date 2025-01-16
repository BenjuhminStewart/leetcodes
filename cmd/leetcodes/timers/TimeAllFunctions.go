package timers

import (
	"fmt"

	"github.com/fatih/color"
)

func TimeAllFunctions(n int) {
	TimeArrayAndHashingFunctions(n)
	TimeTwoPointersFunctions(n)
}

func TimeArrayAndHashingFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nArray and Hashing Functions\n")
	fmt.Printf("%v", title)
	TimeContainsDuplicate(n)
	TimeIsAnagram(n)
	TimeTwoSum(n)
}

func TimeTwoPointersFunctions(n int) {
	title := color.New(color.FgBlue, color.Bold).Add(color.Underline).Sprintf("\nTwo Pointers Functions\n")
	fmt.Printf("%v", title)
	TimeValidPalindrome(n)
}
