func majorityElement(nums []int) int {
	vote, count := 0, 0
	for _, n := range nums { 
		if count == 0 { 
			vote = n
		}
		if vote == n { 
			count++
		} else {
			count--
		}
	}

	return vote
}
