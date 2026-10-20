type pairs struct { 
	temp int;
	idx int;
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []pairs{}

	for i, t := range temperatures { 
		for len(stack) > 0 && t > stack[len(stack) - 1].temp { 
			res[stack[len(stack) - 1].idx] = i - stack[len(stack) - 1].idx
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, pairs{
			temp: t,
			idx: i,
		})
	}

	return res
}
