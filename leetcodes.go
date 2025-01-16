package leetcodes

func ContainsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}
	for _, num := range nums {
		_, isSeen := seen[num]
		if isSeen {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func IsAnagram(s string, t string) bool {
	dict := map[rune]int{}
	for _, r := range s {
		dict[r]++
	}

	for _, r := range t {
		dict[r]--
		if dict[r] < 0 {
			return false
		}
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}

func TwoSum(nums []int, target int) []int {
	dict := map[int]int{}
	for i, num := range nums {
		_, contains_target := dict[target-num]
		if contains_target {
			return []int{dict[target-num], i}
		}
		dict[num] = i
	}
	return []int{-1, -1}
}
