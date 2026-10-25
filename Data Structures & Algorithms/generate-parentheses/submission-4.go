func generateParenthesis(n int) []string {
	res := make([]string, 0)
	stack := make([]string, 0)

	var dfs func(open, close int)
	dfs = func(open, close int) { 
		if open == n && close == n{ 
			res = append(res, strings.Join(stack, ""))
			return
		}
		
		if open < n { 
			stack = append(stack, "(")
			dfs(open+1, close)
			stack = stack[:len(stack) - 1]
		}

		if close < open { 
			stack = append(stack, ")")
			dfs(open, close + 1)
			stack = stack[:len(stack) - 1]
		}
	}
	
	dfs(0, 0)

	return res
}
