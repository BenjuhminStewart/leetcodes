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
