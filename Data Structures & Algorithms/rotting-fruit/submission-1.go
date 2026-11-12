func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{}
	fresh := 0
	time := 0
	for r := 0; r < rows; r++ { 
		for c := 0; c < cols; c++ { 
			if grid[r][c] == 1 { 
				fresh++
			}
			if grid[r][c] == 2 { 
				queue = append(queue, [2]int{r,c})
			}
		}
	}

	dirs := [4][2]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}}
	for fresh > 0 && len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ { 
			val := queue[0]
			queue = queue[1:]
			
			for _, d := range dirs { 
				newRow, newCol := val[0] + d[0], val[1] + d[1]
				if newRow >= 0  && newRow < rows &&  newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 { 
					grid[newRow][newCol] = 2
					fresh--
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
			
		}
		time++


	}
    
	if fresh > 0 { 
		return -1
	}

	return time
}
