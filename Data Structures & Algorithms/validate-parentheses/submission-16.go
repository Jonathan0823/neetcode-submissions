func isValid(s string) bool {
	if len(s) % 2 != 0 { 
		return false
	}
	
    stack := make([]byte, 0, len(s)/2)
	pairs := map[byte]byte{ 
		']':'[',
		'}': '{',
		')': '(',
	}

	for i := 0; i < len(s); i++ { 
		b := s[i]
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
