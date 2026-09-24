func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens { 
		val, err := strconv.Atoi(token)
		if err != nil { 
			first, second := stack[len(stack) - 2], stack[len(stack) - 1]
			stack = stack[:len(stack)-2]

			switch token { 
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
