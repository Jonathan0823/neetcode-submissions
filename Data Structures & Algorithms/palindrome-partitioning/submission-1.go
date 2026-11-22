func partition(s string) [][]string {
	res := [][]string{}
	stack := []string{}

	var dfs func(int)
	dfs = func(i int) { 
		if i == len(s) { 
			res = append(res, append([]string{}, stack...))
			return
		}

		for j := i; j < len(s); j++ { 
			if isPali(s, i, j) { 
				stack = append(stack, s[i:j+1])
				dfs(j+1)
				stack = stack[:len(stack) - 1]
			}
		}
	}

	dfs(0)
	return res
}

func isPali(s string, l, r int) bool { 
	for l < r { 
		if s[l] != s[r] { 
			return false
		}
		l++
		r--
	}
	return true
}
