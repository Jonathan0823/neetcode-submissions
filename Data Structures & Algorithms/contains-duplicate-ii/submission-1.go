func containsNearbyDuplicate(nums []int, k int) bool {
	position := make(map[int]int, len(nums))
	for i, n := range nums { 
		if val, exist := position[n]; exist { 
			dist := i - val
			if dist <= k { 
				return true
			}
		}
		position[n] = i
	}

	return false
}
