func calPoints(operations []string) int {
	stack := []int{}
	for _, o := range operations { 
		switch o { 
			case "+":
				stack = append(stack, stack[len(stack) - 2] + stack[len(stack) - 1])
			case "D":
				stack = append(stack, 2*stack[len(stack) - 1])
			case "C":
				stack = stack[:len(stack) - 1]
			default:
				val, _ := strconv.Atoi(o)
				stack = append(stack, val)
		}
	}

	res := 0
	for _, s := range stack { 
		res += s
	}
	return res
}
