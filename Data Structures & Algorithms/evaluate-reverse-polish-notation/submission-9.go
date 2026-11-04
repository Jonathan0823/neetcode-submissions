func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens { 
		switch t { 
			case "+":
				first := stack[len(stack) - 2]
				second := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, first + second)
			case "-":
				first := stack[len(stack) - 2]
				second := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, first - second)
			case "/":
				first := stack[len(stack) - 2]
				second := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, first / second)
			case "*":
				first := stack[len(stack) - 2]
				second := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				stack = append(stack, first * second)
			default:
				val, _ := strconv.Atoi(t)
				stack = append(stack, val)
		}
	}

	return stack[0]
}
