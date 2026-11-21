func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	dir := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(int, int, int) bool 
	dfs = func(i, r, c int) bool { 
		if i == len(word) { 
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[i] {
			return false
		}

		temp := board[r][c] 
		board[r][c] = '#'
		for _, d := range dir { 
			if dfs(i+1, r+d[0], c+ d[1]) { 
				return true
			}
		}
		board[r][c] = temp

		return false
	}

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if dfs(0, r, c) { 
				return true
			}
		}
	}

	return false
}
