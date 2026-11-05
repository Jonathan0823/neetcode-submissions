func maxProduct(nums []int) int {
   	n := len(nums)
	dpMin := make([]int, n)
	dpMax := make([]int, n)
	dpMin[n - 1] = nums[n - 1]
	dpMax[n - 1] = nums[n - 1]

	ans := nums[n - 1]

	for i := n - 2; i >= 0; i-- { 
		c1 := nums[i]
		c2 := nums[i] * dpMin[i + 1]
		c3 := nums[i] * dpMax[i + 1]

		dpMin[i] = min(c1, min(c2, c3))
		dpMax[i] = max(c1, max(c2, c3))

		ans = max(ans, dpMax[i])
	}

	return ans
}
