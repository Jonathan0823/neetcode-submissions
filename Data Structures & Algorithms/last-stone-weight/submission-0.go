func lastStoneWeight(stones []int) int {
	for len(stones) > 1 { 
		sort.Ints(stones)

		curr := stones[len(stones) - 1] - stones[len(stones) - 2]
		stones = stones[:len(stones)-2] 
		if curr > 0 { 
			stones = append(stones, curr)
		}
	}

	return stones[0]
}
