func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)

	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) { 
		coord := [2]int{r, c}
		if visit[coord] || r < 0 || r >= rows || c < 0 || c >= cols || heights[r][c] < prevHeight { 
			return
		}

		visit[coord] = true

		dfs(r + 1, c, visit, heights[r][c])
		dfs(r - 1, c, visit, heights[r][c])
		dfs(r, c + 1, visit, heights[r][c])
		dfs(r, c - 1, visit, heights[r][c])
	}

	for c := 0; c < cols; c++ { 
		dfs(0, c, pac, heights[0][c])
		dfs(rows - 1, c, atl, heights[rows - 1][c])
	}

	for r := 0; r < rows; r++ { 
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols - 1, atl, heights[r][cols - 1])
	}

	res := make([][]int, 0)
	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			coord := [2]int{r, c}
			if pac[coord] && atl[coord] { 
				res = append(res, []int{r,c})
			}
		}
	}

	return res
}
