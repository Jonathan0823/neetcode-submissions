func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	
	for i, n := range nums { 
		complement := target - n
		if compIdx, exist := seen[complement]; exist { 
			return []int{compIdx, i}
		}
		seen[n] = i
	}

	return []int{}
}
