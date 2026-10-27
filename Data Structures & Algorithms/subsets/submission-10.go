func subsets(nums []int) [][]int {
	var res [][]int
	var sets []int

	var dfs func(i int)
	dfs = func(i int) { 
		if i == len(nums) {
			res = append(res, append([]int{}, sets...))
			return
		}

		sets = append(sets, nums[i])
		dfs(i + 1)
		sets = sets[:len(sets) - 1]
		dfs(i + 1)
	}
	dfs(0)

	return res
}
