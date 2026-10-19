func calPoints(operations []string) int {
	stack := []int{}
	for _, o := range operations {
		switch o { 
			case "+":
				stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
			case "C":
				stack = stack[:len(stack) - 1]
			case "D":
				stack = append(stack, stack[len(stack) - 1] * 2)
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
