package lc36

//isValidSudoku https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	// '.' -> 46, '1' -> 49, '9' -> '57'
	m, n := 9, 9
	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	subs := make(map[int]map[byte]bool)
	for i := 0; i < m; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		subs[i] = make(map[byte]bool)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			if rows[i][b] {
				return false
			} else {
				rows[i][b] = true
			}

			if cols[j][b] {
				return false
			} else {
				cols[j][b] = true
			}

			k := i/3*3 + j/3
			if subs[k][b] {
				return false
			} else {
				subs[k][b] = true
			}
		}
	}
	return true
}
