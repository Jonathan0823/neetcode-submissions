func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens { 
		res := 0
		if val, err := strconv.Atoi(t); err != nil {
			first := stack[len(stack) - 2]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			switch t { 
				case "+":
					res = first + second
				case "-":
					res = first - second
				case "*":
					res = first * second
				case "/":
					res = first / second
			}
		} else { 
			res = val
		}
		stack = append(stack, res)
	}

	return stack[0]
}
