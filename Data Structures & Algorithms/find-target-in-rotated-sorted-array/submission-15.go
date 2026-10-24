func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r { 
		m := l + (r - l)/2
		if target == nums[m] { 
			return m
		}

		if nums[m] >= nums[l] { 
			if target < nums[l] || target > nums[m] { 
				l = m + 1
			} else { 
				r = m - 1
			}
		} else {
			if target > nums[r] || target < nums[m]{ 
				r = m - 1
			} else { 
				l = m + 1
			} 
		}
	}

	return -1
}
