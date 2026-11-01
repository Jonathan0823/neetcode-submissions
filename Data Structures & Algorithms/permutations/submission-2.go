func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(int) 

	backtrack = func(first int) { 
		if first == len(nums) { 
			res = append(res, append([]int{}, nums...))
			return
		}

		for i := first; i < len(nums); i++ { 
			nums[i], nums[first] = nums[first], nums[i]
			backtrack(first+1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	backtrack(0)

	return res

}
