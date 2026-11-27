func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for _, n := range nums { 
		numMap[n] = true
	}

	res := 0
	for _, n := range nums { 
		count := 1
		// start of sequence
		if !numMap[n - 1] { 
			for numMap[n + count] { 
				count++
			}

			res = max(res, count)
		}
	}

	return res
}
