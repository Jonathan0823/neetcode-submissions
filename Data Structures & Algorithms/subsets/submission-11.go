func subsets(nums []int) [][]int {
	var res [][]int
	stack := []int{}

	var dfs func(int) 
	dfs = func(i int) { 
		if i == len(nums) { 
			res = append(res, append([]int{}, stack...))
			return
		}

		stack = append(stack, nums[i])
		dfs(i+1)
		stack = stack[:len(stack) - 1]
		dfs(i+1)
	}

	dfs(0)
	return res
}
