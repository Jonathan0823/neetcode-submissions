func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(int)
	dfs = func(i int) { 
		if i >= len(nums) { 
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset) - 1]
		dfs(i + 1)
	}

	dfs(0)

	return res
}



