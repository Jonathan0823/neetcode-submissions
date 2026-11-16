func permute(nums []int) [][]int {
	var res [][]int

	var dfs func(int)
	dfs = func(first int) { 
		if first == len(nums) { 
			res = append(res, append([]int{}, nums...))
			return
		}

		for i := first; i < len(nums); i++ { 
			nums[i], nums[first] = nums[first], nums[i]
			dfs(first+1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	dfs(0)

	return res
}
