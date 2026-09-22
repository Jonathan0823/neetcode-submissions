func longestConsecutive(nums []int) int {
	res := 0
	numSet := make(map[int]bool, len(nums))

	for _, num := range nums { 
		numSet[num] = true
	}

	for _, num := range nums { 
		// check if its the start of sequence
		if !numSet[num - 1] { 
			//start of sequence
			length := 1
			for numSet[num +length]{
					length++
				} 
			
			res = max(res, length)
		}
	}

	return res
}
