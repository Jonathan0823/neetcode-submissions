func longestConsecutive(nums []int) int {
	numsSets := make(map[int]struct{})
	for _, num := range nums { 
		numsSets[num] = struct{}{}
	}

	longest := 0

	for _, num := range nums { 
		if _, exist := numsSets[num - 1]; !exist { 
			// it's start of the sequence
			length := 1

			// loop the sequence
			for { 
				if _, exist := numsSets[num + length]; exist { 
					length++
				} else { 
					break
				}
			} 

			longest = max(longest, length)
		}
	}

	return longest

}
