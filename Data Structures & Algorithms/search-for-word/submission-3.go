func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	path := make(map[[2]int]bool)

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool { 
		if i == len(word) { 
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || 
		board[r][c] != word[i] || path[[2]int{r, c}] { 
			return false
		}
		
		path[[2]int{r, c}] = true
		res := dfs(r + 1, c, i+1) ||
			dfs(r - 1, c, i+1) ||
			dfs(r, c + 1, i+1) ||
			dfs(r, c - 1, i+1)
		delete(path, [2]int{r, c})

		return res
	}

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if dfs(r, c, 0) { 
				return true
			}
		}
	}

	return false
}
