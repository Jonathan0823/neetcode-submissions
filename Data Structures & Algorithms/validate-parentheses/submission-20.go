func isValid(s string) bool {
	if len(s) % 2 != 0 { 
		return false
	}

    stack := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s); i++ { 
		b := s[i]
		
		switch b {
			case ']', ')', '}':
				if len(stack) < 1 { 
					return false
				}
				val := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]	

				switch b { 
					case ']':
						if val != '[' { 
							return false
						}
					case ')':
						if val != '(' { 
							return false
						}
					case '}':
						if val != '{' { 
							return false
						}
				}
			default:
				stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
