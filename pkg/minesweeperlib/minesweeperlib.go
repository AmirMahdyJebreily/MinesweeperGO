package minesweeperlib

func GetCellNumbers(rows, cols int, bombs map[int]int) [][]int {
	board := make([][]int, cols)

	// init the board
	for i := range board {
		board[i] = make([]int, rows)
	}

	for x, y := range bombs {
		// add one score to 8 cells around bombs
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				u, v := x+i, y+j
				if (u < cols && u >= 0) && (v < rows && v >= 0) {
					
					board[u][v]++
				}
			}
		}

	}
	return board
}
