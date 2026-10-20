func missingNumber(nums []int) int {
	xor := len(nums)
	for idx, n := range nums { 
		xor ^= idx ^ n
	}

	return xor
}
