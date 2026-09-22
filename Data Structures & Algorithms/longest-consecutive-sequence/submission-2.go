func longestConsecutive(nums []int) int {
	res := 0
	numSet := make(map[int]struct{})

	for _, num := range nums { 
		numSet[num] = struct{}{}
	}

	for _, num := range nums { 
		// check if its the start of sequence
		if _, exist := numSet[num - 1]; !exist { 
			//start of sequence
			length := 1
			for {
				if _, exist := numSet[num +length]; exist { 
					length++
				} else { 
					break
				}
			}
			res = max(res, length)
		}
	}

	return res
}
