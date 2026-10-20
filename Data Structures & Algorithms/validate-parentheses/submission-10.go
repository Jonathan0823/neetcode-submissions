func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune { 
		']':'[',
		')':'(',
		'}':'{',
	}

	for _, r := range s { 
		if pair, exist := pair[r]; exist { 
			if len(stack) < 1 { 
				return false
			}

			last := stack[len(stack) - 1]
			if last != pair { 
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		} else { 
			stack = append(stack, r)
		}
	}
	

	return len(stack) == 0
    
}
