func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int

	sort.Ints(candidates)

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) { 
		if total  == target { 
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if total > target || i == len(candidates)  { 
			return 
		}

		cur = append(cur, candidates[i])
		dfs(i+1, cur, total + candidates[i])
		cur = cur[:len(cur) - 1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] { 
			i++
		}
		dfs(i+1, cur, total)
	}

	dfs(0, []int{}, 0)
	return res
}
