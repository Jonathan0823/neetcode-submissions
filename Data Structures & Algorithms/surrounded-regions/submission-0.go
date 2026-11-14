func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	directions := [4][2]int{{1,0}, {-1,0}, {0, 1}, {0, -1}}

	var dfs func(r, c int)
	dfs = func(r, c int) { 
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' { 
			return
		}

		board[r][c] = 'T'
		for _, d := range directions { 
			dfs(r + d[0], c + d[1])
		}
	}

	for c := 0; c < cols; c++ { 
		dfs(0, c)
		dfs(rows - 1, c)
	}
    
	for r := 0; r < rows; r++ { 
		dfs(r, 0)
		dfs(r, cols - 1)
	}

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if board[r][c] == 'O' { 
				board[r][c] = 'X'
			}
		}
	}

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if board[r][c] == 'T' { 
				board[r][c] = 'O'
			}
		}
	}
}
