func combinationSum(nums []int, target int) [][]int {
	var res [][]int

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) { 
		if total == target { 
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if total > target || i >= len(nums) { 
			return
		}
		
		cur = append(cur, nums[i])
		dfs(i, cur, total + nums[i])
		cur = cur[:len(cur) - 1]
		dfs(i + 1, cur, total)
	}

	dfs(0, []int{}, 0)
	return res
}
