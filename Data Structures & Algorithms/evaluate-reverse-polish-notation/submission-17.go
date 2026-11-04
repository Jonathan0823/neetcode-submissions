func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens)/2+1)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, b+a)
			case "-":
				stack = append(stack, b-a)
			case "*":
				stack = append(stack, b*a)
			case "/":
				stack = append(stack, b/a)
			}
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}