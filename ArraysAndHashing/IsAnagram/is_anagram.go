package leetcodes

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
