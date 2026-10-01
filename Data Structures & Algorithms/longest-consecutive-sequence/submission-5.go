func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	maxSeq := 0

	for _, num := range nums { 
		seen[num] = true
	}

	for _, num := range nums { 
		// start of the sequence
		if !seen[num - 1] { 
			count := 1

			for seen[num + count] { 
				count++
			}

			maxSeq = max(maxSeq, count)

		}
	}

	return maxSeq
}
