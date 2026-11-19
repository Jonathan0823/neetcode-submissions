func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)

	var dfs func(int, int)
	dfs = func(openN, closedN int) { 
		if openN == n && closedN == n { 
			res = append(res, strings.Join(stack,""))
			return
		}

		if openN < n { 
			stack = append(stack, "(")
			dfs(openN + 1, closedN)
			stack = stack[:len(stack) - 1]
		}

		if closedN < openN { 
			stack = append(stack, ")")
			dfs(openN, closedN + 1)
			stack = stack[:len(stack) - 1]
		}
	}

	dfs(0, 0)

	return res
}
