package leetcodes

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
