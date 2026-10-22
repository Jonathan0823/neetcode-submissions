func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, len(position))
	stack := []float64{}
	for i, p := range position { 
		pairs[i] = [2]int{p, speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool { 
		return pairs[i][0] > pairs[j][0]
	})

	for _, p := range pairs { 
		time := float64(target - p[0]) / float64(p[1])
		if len(stack) == 0 || time > stack[len(stack) - 1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}
