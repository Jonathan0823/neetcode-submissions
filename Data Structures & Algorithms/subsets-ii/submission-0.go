func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	subset := []int{}

	var dfs func(i int)
	dfs = func(i int) { 
		if i >= len(nums) { 
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		subset = append(subset, nums[i])
		dfs(i+1)

		subset = subset[:len(subset) - 1]
		for i+1 < len(nums) && nums[i] == nums[i+1] { 
			i++
		}
		dfs(i+1)
	}
	dfs(0)

	return res
}
