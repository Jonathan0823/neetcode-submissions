func majorityElement(nums []int) int {
	res, count := 0, 0

	for _, n := range nums { 
		if count == 0  {
			res = n 
		}

		if n != res { 
			count--
		} else { 
			count++
		}
	}

	return res
    
}
