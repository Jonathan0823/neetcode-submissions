func lastStoneWeight(stones []int) int {
	for len(stones) > 1 { 
		sort.Ints(stones)

		curr := stones[len(stones) - 1] - stones[len(stones) - 2]
		stones = stones[:len(stones)-2] 
		if curr > 0 { 
			stones = append(stones, curr)
		}
	}

	if (len(stones) == 0) { 
		return 0
	}

	return stones[0]
}
