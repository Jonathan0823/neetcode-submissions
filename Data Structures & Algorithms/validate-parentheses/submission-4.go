func isValid(s string) bool {
	stack := []rune{}
	closeToOpen :=  map[rune]rune{
		'}':'{',
		']':'[',
		')':'(',
	}

	for _, b := range s { 
		if open, exist := closeToOpen[b]; exist { 
			if len(stack) > 0 {
				top := stack[len(stack) - 1]
				stack = stack[:len(stack)  - 1] 
				if top != open { 
					return false
				}

			} else { 
				return false
			}

		} else { 
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
