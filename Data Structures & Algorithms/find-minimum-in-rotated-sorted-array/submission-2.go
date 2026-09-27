func findMin(nums []int) int {
	res := nums[0]

	l, r := 0, len(nums) - 1
	for l <= r { 
		if nums[l] < nums[r] { 
			res = min(res, nums[l])
			break
		} 

		m := l + (r - l)/2
		res = min(res, nums[m])
		if nums[m] >= nums[l] { 
			l = m +1
		} else { 
			r = m - 1
		}
		
	}

	return res

}
