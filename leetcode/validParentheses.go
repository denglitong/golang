package leetcode

func isPair(c1, c2 byte) bool {
	return (c1 == '(' && c2 == ')') ||
		(c1 == '{' && c2 == '}') ||
		(c1 == '[' && c2 == ']')
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	var stack []byte
	for _, c := range s {
		if len(stack) > 0 && isPair(stack[len(stack)-1], byte(c)) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	return len(stack) == 0
}
