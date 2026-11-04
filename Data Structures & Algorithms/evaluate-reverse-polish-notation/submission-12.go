func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens { 
		if val, err := strconv.Atoi(t); err == nil {
			stack = append(stack, val)
			continue
		}

		first := stack[len(stack) - 2]
		second := stack[len(stack) - 1]
		stack = stack[:len(stack) - 2]
		res := 0

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
		stack = append(stack, res)
	}

	return stack[0]
}
