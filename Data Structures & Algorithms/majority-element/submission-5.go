func majorityElement(nums []int) int {
	count, majority := 0, 0
	for _, n := range nums { 
		if count == 0 { 
			majority = n
		}

		if n == majority { 
			count++
		} else { 
			count--
		}
	}

	return majority
}
