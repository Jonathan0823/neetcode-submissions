func canPartition(nums []int) bool {
	target := 0
	for _, n := range nums { 
		target += n
	}
	if target % 2 != 0 { 
		 return false
	}
	target /= 2
    dp := make([]bool, target + 1) // idx 0 - 5
	dp[0] = true // base case

	for _, n := range nums {
		for t := target; t >= n; t-- { 
			dp[t] = dp[t] || dp[t - n]
		}
	}

	return dp[target]
}
