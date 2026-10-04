func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens { 
		if val, err := strconv.Atoi(t); err != nil { 
			first := stack[len(stack) - 2]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			switch t {
				case "+":
					stack = append(stack, first + second)
				case "-":
					stack = append(stack, first - second)
				case "*":
					stack = append(stack, first * second)
				case "/":
					stack = append(stack, first / second)
			}
		} else { 
			stack = append(stack, val)
		}
	}

	return stack[0]
}
