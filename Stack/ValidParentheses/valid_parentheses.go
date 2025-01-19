package leetcodes

var (
	allowed = map[rune]bool{
		'(': true,
		')': true,
		'{': true,
		'}': true,
		'[': true,
		']': true,
	}
)

func IsValid(s string) bool {
	if !isClean(s) {
		return false
	}
	parenthesisMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0)
	for _, char := range s {
		_, containsChar := parenthesisMap[char]
		if containsChar {
			stack = append(stack, char)
		} else if len(stack) > 0 && char == parenthesisMap[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func isClean(s string) bool {
	for _, char := range s {
		if _, ok := allowed[char]; !ok {
			return false
		}
	}
	return true
}
