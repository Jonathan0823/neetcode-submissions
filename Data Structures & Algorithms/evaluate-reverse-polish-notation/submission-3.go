func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens { 
		val, err := strconv.Atoi(token)

		if err != nil {
			a :=  stack[len(stack) - 2]
			b := stack[len(stack) - 1]
			stack = stack[:len(stack)-2]
			var res int

			switch token { 
				case "+":
				res = a + b

				case "-":
				res = a -b 
				
				case "*":
				res = a * b
				
				case "/":
				res = a / b
			
			}
			stack = append(stack, res)
		

			continue 
		}

	stack = append(stack, val)
		
	}

	return stack[0]

}
