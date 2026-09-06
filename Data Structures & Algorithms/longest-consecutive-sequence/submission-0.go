func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums { 
		numSet[num] = struct{}{}
	}

	longest := 0
	for num := range numSet { 
		if _, found := numSet[num - 1]; !found { 
			length := 1
			for { 
				if _, exist := numSet[num + length]; exist { 
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
