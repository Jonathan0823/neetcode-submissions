func removeDuplicates(nums []int) int {
	k := 0
	for i, n := range nums { 
		if i > 0 && n == nums[i-1] { 
			continue
		} else { 
			nums[k] = n
			k++
		}
	}

	
	return k
}
