package leetcodes

func GroupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for _, str := range strs {
		chars := make([]rune, 26)
		for _, c := range []rune(str) {
			chars[int(c)-'a']++
		}

		key := string(chars)
		_, ok := dict[key]
		if !ok {
			dict[key] = make([]string, 0)
		}
		dict[key] = append(dict[key], str)
	}

	var result [][]string
	for _, v := range dict {
		result = append(result, v)
	}
	return result
}
