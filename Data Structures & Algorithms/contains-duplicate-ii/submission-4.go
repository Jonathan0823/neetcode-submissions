func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int, len(nums)) // num:idx
	for i, n := range nums { 
		if j, ok := numMap[n]; ok { 
			if (i - j) <= k { 
				return true
			}
		}
		numMap[n] = i
	}

	return false
}
