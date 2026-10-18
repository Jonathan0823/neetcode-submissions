func getConcatenation(nums []int) []int {
    res := make([]int, 2*len(nums))
	for i, n := range nums { 
		res[i] = n
		res[i + len(nums)] = n
	}

	return res
}
