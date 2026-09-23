func searchMatrix(matrix [][]int, target int) bool {
	rows, cols  := len(matrix), len(matrix[0])

	top, bot := 0, rows - 1
	for top <= bot { 
		row := top + (bot - top)/2
		if target > matrix[row][cols-1] { 
			  top = row+1
		} else if target < matrix[row][0] { 
			bot = row - 1
		} else { 
			top = row
			break
		}
	}

	if top > bot {
		return false
	}

	row := top
	l, r := 0, cols -1
	for l <= r { 
		m := l + (r-l)/2
		if target > matrix[row][m] { 
			l = m + 1
		} else if target < matrix[row][m] { 
			r = m - 1
		} else { 
			return true
		}
	}

	return false
}
