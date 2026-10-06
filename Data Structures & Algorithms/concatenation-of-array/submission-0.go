func getConcatenation(nums []int) []int {
	for _, num := range nums { 
		nums = append(nums, num)
	}

	return nums
}
