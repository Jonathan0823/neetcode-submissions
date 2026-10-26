func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, n := range nums { 
		completion := target - n
		if prevIdx, exist := numMap[completion]; exist { 
			return []int{prevIdx, i}
		}
		numMap[n] = i
	}
	return nil
}
