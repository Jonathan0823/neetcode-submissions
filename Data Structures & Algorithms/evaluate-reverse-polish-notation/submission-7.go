func evalRPN(tokens []string) int {
	var stack []int

	for _, t := range tokens { 
		if val, err := strconv.Atoi(t); err != nil {
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
		} else {
			stack = append(stack, val)
		}
	}

	return stack[0]

}
