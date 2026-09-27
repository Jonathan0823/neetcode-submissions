func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l <= r { 
		m := l + (r - l)/2

		if nums[m] == target { 
			return m
		}

		// left sorted portion
		if nums[l] <= nums[m] { 
			if target > nums[m] || target < nums[l]{ 
				l = m + 1
			} else { 
				r = m - 1
			}
		} else { 
			// right sorted
			if target < nums[m] || target > nums[r] { 
				r = m - 1
			} else { 
				l = m + 1
			}
			
		}
	}

	return -1

}
