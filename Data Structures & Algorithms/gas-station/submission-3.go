func canCompleteCircuit(gas []int, cost []int) int {
    if sum(gas) < sum(cost) { 
		return -1
	}

	total := 0
	res := 0
	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 { 
			total = 0
			res = i + 1
		}
	}

	return res
}

func sum(nums []int) int { 
	var total int
	for _, n := range nums { 
		total += n
	}

	return total
}
