func lengthOfLIS(nums []int) int {
	LIS := make([]int, len(nums))
	for i := range LIS { 
		LIS[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- { 
		for j := i + 1; j < len(nums); j++ { 
			if nums[i] < nums[j] { 
				if 1 + LIS[j] > LIS[i] { 
					LIS[i] = 1 + LIS[j]
				}
			}
		}
	}

	maxLen := 1
	for _, len := range LIS { 
		maxLen = max(maxLen, len )
	}

	return maxLen    
}
