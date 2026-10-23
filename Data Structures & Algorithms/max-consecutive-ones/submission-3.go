func findMaxConsecutiveOnes(nums []int) int {
	res, cnt := 0, 0
	for _, n := range nums { 
		if n != 1 { 
			res = max(res, cnt)
			cnt = 0
		} else { 
			cnt++
		}
	}

	return max(res, cnt)	
}
