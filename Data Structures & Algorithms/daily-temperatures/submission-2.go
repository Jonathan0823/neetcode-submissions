func dailyTemperatures(temperatures []int) []int {
	res := make([]int,len(temperatures))
	stack := []int{}

	for i, t := range temperatures { 
		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < t { 
			res[stack[len(stack) - 1]] = i - stack[len(stack) - 1] 
			stack = stack[:len(stack) - 1]
		} 

		stack = append(stack, i)

	}

	return res
}
