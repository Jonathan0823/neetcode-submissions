func missingNumber(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res ^= n ^ i
	}

	return res
}
