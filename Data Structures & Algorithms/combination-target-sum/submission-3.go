func combinationSum(nums []int, target int) [][]int {
	stack := []int{}
	var res [][]int

	var dfs func(i, sum int)
	dfs = func(i, sum int) { 
		if sum == target { 
			res = append(res, append([]int{}, stack...))
			return
		}
		if i >= len(nums) || sum > target  { 
			return 
		}

		stack = append(stack, nums[i])
		dfs(i, sum + nums[i])
		stack = stack[:len(stack) - 1]
		dfs(i + 1, sum)
	}

	dfs(0, 0)

	return res
}
