func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	// prefix
	for i := 1; i < len(res); i++ { 
		res[i] = res[i-1] * nums[i-1]
	}

	//suffix
	right := 1
	for i := len(nums) - 1; i >= 0; i-- { 
		res[i] *= right
		right *= nums[i]
	}

	return res
}
