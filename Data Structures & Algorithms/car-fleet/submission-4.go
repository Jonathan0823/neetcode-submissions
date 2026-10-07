func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, len(position))
	for i := 0; i < len(position);i++ { 
		pairs[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool { 
		return pairs[i][0] > pairs[j][0]
	})

	stack := []float64{}
	for _, p := range pairs { 
		time := float64(target - p[0])/ float64(p[1])
		if len(stack) == 0 || time > stack[len(stack) - 1]  { 
			stack = append(stack, time)
		}

	}

	return len(stack)
}
