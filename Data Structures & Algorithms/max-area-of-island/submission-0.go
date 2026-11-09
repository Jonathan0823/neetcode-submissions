func maxAreaOfIsland(grid [][]int) int {
	directions := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) int 
	dfs = func(r, c int) int { 
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 { 
			return 0
		}
		area := 0

		grid[r][c] = 0
		area++
		for _, d := range directions { 
			area += dfs(r + d[0], c + d[1])
		}
		return area
	}

	res := 0

	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			res = max(res, dfs(r,c))
		}
	}

	return res
}
