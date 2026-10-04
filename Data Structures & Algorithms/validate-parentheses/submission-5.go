func isValid(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune { 
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, r := range s { 
		if pair, exist := pairs[r]; exist { 
			if len(stack) < 1 { 
				return false
			}

			if pair == stack[len(stack) - 1] { 
				stack = stack[:len(stack) - 1]
			} else { 
				return false
			}
			continue
		}
		stack = append(stack, r)
	}

	return len(stack) == 0
}
