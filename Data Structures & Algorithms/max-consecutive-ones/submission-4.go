func findMaxConsecutiveOnes(nums []int) int {
	res, cnt := 0, 0
	for _, n := range nums { 
		if n == 1 { 
			cnt++
			res = max(res, cnt)
		} else { 
			cnt = 0
		}
	}

	return res
}
