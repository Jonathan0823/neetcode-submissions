func missingNumber(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums  {
		numSet[num] = struct{}{}
	}

	for i := 0; i < len(nums); i++ { 
		if _, exist := numSet[i]; !exist { 
			return i
		}
	}

	return len(nums) 

}
