func missingNumber(nums []int) int {
	n := len(nums)
	xor := n
	for i := 0; i< n; i++ { 
		xor  ^= i ^ nums[i]
	}

	return xor

}
