func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix) - 1

	row := -1
	for l <= r { 
		mid := l + (r-l)/2

		rowFirst, rowLast := matrix[mid][0], matrix[mid][len(matrix[mid]) - 1]
		if target >= rowFirst && target <= rowLast {
			row = mid
			break
		} else if target < rowFirst { 
			r = mid-1
		} else if target > rowLast { 
			l = mid + 1
		}
	}	

	 if !(l <= r) {
        return false
    }

	l, r = 0, len(matrix[0]) - 1
	for l <= r { 
		mid := l + (r - l)/2
		if matrix[row][mid] == target { 
			return true
		} else if matrix[row][mid] < target { 
			l = mid + 1
		} else if matrix[row][mid] > target{ 
			r = mid - 1
		}
	}

	return false

}
