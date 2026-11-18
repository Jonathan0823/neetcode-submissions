func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	stack := []int{}

	var dfs func(int)
	dfs = func(i int) { 
		if i >= len(nums) { 
			res = append(res, append([]int{}, stack...))
			return
		}

		stack = append(stack, nums[i])
		dfs(i + 1)
		for i + 1 < len(nums) && nums[i] == nums[i+1] { 
			i++
		}
		stack = stack[:len(stack) - 1]
		dfs(i+1)
	}
	dfs(0)

	return res
}
