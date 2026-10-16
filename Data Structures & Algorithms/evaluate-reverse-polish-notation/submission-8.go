func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, t := range tokens { 
		val, err := strconv.Atoi(t); 
		if err != nil {
			res := 0
			first, sec := stack[len(stack) - 2], stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			switch t { 
				case "+":
					res = first + sec
				case "-":
					res = first - sec
				case "*":
					res = first * sec
				case "/":
					res = first / sec
			}

			stack = append(stack, res)
			continue
		} 		
		stack = append(stack, val)
		
	}

	return stack[0]

}
