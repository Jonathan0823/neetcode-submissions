func getConcatenation(nums []int) []int {
	length := len(nums)
	res := make([]int, length * 2)
	for i, n := range nums { 
		res[i] = n
		res[i + length] = n
	}

	return res
}
