func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var dfs func(open, close int, cur string)
	dfs = func(open, close int, cur string) { 
		if len(cur) == n * 2 { 
			res = append(res, cur)
			return
		}
		
		if open < n { 
			dfs(open+1, close, cur + "(")
		}

		if close < open { 
			dfs(open, close + 1, cur + ")")
		}
	}
	
	dfs(0, 0, "")

	return res
}
