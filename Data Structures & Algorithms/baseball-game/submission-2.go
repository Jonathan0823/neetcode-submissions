func calPoints(operations []string) int {
	stack := []int{}
	for _, o := range operations {
		if val, err := strconv.Atoi(o); err == nil { 
			stack = append(stack, val)
		} else { 
			switch o { 
				case "+":
					stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
				case "C":
					stack = stack[:len(stack) - 1]
				case "D":
					stack = append(stack, stack[len(stack) - 1] * 2)
			}
		}
	}

	res := 0
	for _, s := range stack { 
		res += s
	}

	return res
}
