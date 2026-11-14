func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	stack := []int{}
	
	var dfs func(i, sum int)
	dfs = func(i, sum int) { 
		if sum == target { 
			res = append(res, append([]int{}, stack...))
			return
		}
		
		if i == len(candidates) || sum > target { 
			return
		}

		stack = append(stack, candidates[i])
		dfs(i + 1, sum + candidates[i])
		stack = stack[:len(stack) - 1]
		for i+1 < len(candidates) && candidates[i] == candidates[i + 1] { 
			i++
		}
		dfs(i + 1, sum)
	}

	dfs(0, 0)
	return res
}
