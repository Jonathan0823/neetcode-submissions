func numIslands(grid [][]byte) int {
	directions := [][]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}}
    rows, cols := len(grid), len(grid[0])
	islands := 0

	var dfs func(r, c int)
	dfs = func(r, c int) { 
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' { 
			return
		}

		grid[r][c] = '0'
		for _, dir := range directions { 
			dfs(r+dir[0], c+dir[1])
		}

	}

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if grid[r][c] == '1' { 
				dfs(r,c)
				islands++
			}
		}
	}

	return islands
}
