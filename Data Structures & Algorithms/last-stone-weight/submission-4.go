func lastStoneWeight(stones []int) int {
	for len(stones) > 1 { 
		sort.Ints(stones)

		smashed := stones[len(stones) - 1] - stones[len(stones) - 2]
		stones = stones[:len(stones) - 2]
		stones = append(stones, smashed)
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}
