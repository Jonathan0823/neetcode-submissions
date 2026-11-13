func isValid(s string) bool {
    stack := []rune{}
	pairs := map[rune]rune{ 
		']':'[',
		'}': '{',
		')': '(',
	}

	for _, b := range s { 
		if pair, found := pairs[b]; found { 
			if len(stack) < 1 { 
				return false
			}

			val := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if val != pair { 
				return false	
			}
		} else { 
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
