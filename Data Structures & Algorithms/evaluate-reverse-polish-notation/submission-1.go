func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens { 
		val, err := strconv.Atoi(token)
		if err != nil {
			switch token { 
			case "+":
				res := int(stack[len(stack) - 1]) + int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "-":
			res := int(stack[len(stack) - 1]) - int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "*":
			res := int(stack[len(stack) - 1]) * int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			case "/":
			res := int(stack[len(stack) - 1]) / int(stack[len(stack) - 2])
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			
		}
			continue 
		}

	stack = append(stack, val)
		
	}

	return stack[0]

}
