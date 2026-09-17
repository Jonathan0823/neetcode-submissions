func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	pair := make([][2]int, n)
	for i := 0; i < n; i++ { 
		pair[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pair, func(i, j int) bool { 
		return pair[i][0] > pair[j][0]
	})

	stack := []float64{}
	for _, p := range pair { 
		time := float64(target-p[0]) / float64(p[1])
		stack = append(stack, time)
		if len(stack) >1 && stack[len(stack)-1] <= stack[len(stack)-2] { 
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)

}
