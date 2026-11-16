func findMaxConsecutiveOnes(nums []int) int {
	count, res := 0, 0
	for _, n := range nums { 
		if n == 0 { 
			res = max(res, count)
			count = 0
			continue
		}
		count++
	}

	return max(count, res)
}
