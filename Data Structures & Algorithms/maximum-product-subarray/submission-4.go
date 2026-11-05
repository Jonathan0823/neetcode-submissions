func maxProduct(nums []int) int {
   	n := len(nums)
	currMax := nums[n- 1]
	currMin := nums[n- 1]

	ans := nums[n - 1]

	for i := n - 2; i >= 0; i-- { 
		prevMax := currMax
		prevMin := currMin
		c1 := nums[i]
		c2 := nums[i] * prevMin
		c3 := nums[i] * prevMax

		currMin = min(c1, min(c2, c3))
		currMax = max(c1, max(c2, c3))

		ans = max(ans, currMax)
	}

	return ans
}
