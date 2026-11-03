func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	l, r := 0, rows * cols - 1
	for l <= r { 
		m := l + (r - l)/ 2
		if matrix[m/cols][m%cols] == target { 
			return true
		} else if matrix[m/cols][m%cols] < target { 
			l = m + 1
		} else { 
			r = m - 1
		}
		
	}

	return false
}
