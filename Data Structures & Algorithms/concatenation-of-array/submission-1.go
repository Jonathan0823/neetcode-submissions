func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for i, nums := range nums { 
		ans[i] = nums
		ans[i+n] = nums
	}

	return ans
}
