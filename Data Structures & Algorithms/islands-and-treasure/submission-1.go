func islandsAndTreasure(grid [][]int) {
	rows, cols := len(grid), len(grid[0])
    q := [][2]int{}

	for r := 0; r < rows; r++ { 
		for c := 0;  c < cols; c++ { 
			if grid[r][c] == 0 { 
				q = append(q, [2]int{r, c})
			}
		}
	}

	dir := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 { 
		val := q[0]
		q = q[1:]
		r, c := val[0], val[1]

		for _, d := range dir { 
			nr, nc := r + d[0], c + d[1]
			
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 2147483647 { 
				grid[nr][nc] = grid[r][c] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}
	return 
}
