package algorithm

func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ']' && top != '[' {
				return false
			}
			if s[i] == '}' && top != '{' {
				return false
			}
			if s[i] == ')' && top != '(' {
				return false
			}

		}
	}

	return len(stack) == 0
}

func isValid2(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	var stack []byte
	var mapping = map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	for i := 0; i < n; i++ {
		if val, ok := mapping[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
