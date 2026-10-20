func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, n := range nums { 
		numSet[n] = true
	}

	res := 0
	for _, n := range nums { 
		if !numSet[n-1] { 
			cnt := 1
			for numSet[n+cnt] { 
				cnt++
			}

			res = max(res, cnt)
		}
	}

	return res
}
