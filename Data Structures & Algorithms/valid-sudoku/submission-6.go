func isValidSudoku(board [][]byte) bool {
	var rows, cols, square [9][9]bool

	for r := 0; r < 9; r++ { 
		for c := 0; c < 9; c++ { 
			if board[r][c] == '.' { 
				continue
			}

			val := board[r][c]
			idx := val - '1'
			squareIdx := (r / 3) * 3 + (c / 3)
			if rows[r][idx] || cols[c][idx] || square[squareIdx][idx] { 
				return false
			}

			rows[r][idx] = true
			cols[c][idx] = true
			square[squareIdx][idx] = true

		}
	}

	return true

}