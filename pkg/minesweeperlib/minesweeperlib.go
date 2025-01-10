package minesweeperlib

import (
	"math"
	"math/rand"
	"slices"
)

func GetBoard(cols, rows int) *[][]int {
	board := make([][]int, rows)

	// init the board
	for i := range board {
		board[i] = make([]int, cols)
	}

	return &board
}
func GetRandomBombs(cols, rows, x0, y0 int, count int) *[][2]int {

	bombs := make([][2]int, 0)

	// The radius of the circle that determines how many points around the point (x0, y0) are forbidden points.
	const radiusFactor = 0.16            // TODO:Add the degree of hardness: the harder, the smaller the radius
	var exclusionCenter = [2]int{x0, y0} // the point that user clicked on it for first (x0 , y0)

	// The radius of the circle must be specified based on the width of the page.
	calculateExclusionRadius := func() float64 {
		return radiusFactor * float64(cols)
	}

	isInExclusionZone := func(x, y int) bool {
		radius := calculateExclusionRadius()
		distance := math.Sqrt(math.Pow(float64(x-exclusionCenter[0]), 2) + math.Pow(float64(y-exclusionCenter[1]), 2))
		return distance < radius
	}

	isFree := func(x, y int) bool {
		return x >= 0 && x < cols && y >= 0 && y < rows && !slices.Contains(bombs, [2]int{x, y}) && !isInExclusionZone(x, y)
	}

	attachParticle := func(x, y int) {
		bombs = append(bombs, [2]int{x, y})
	}

	findRandomPoint := func() (int, int) {
		for {
			x, y := rand.Intn(cols), rand.Intn(rows)
			if isFree(x, y) {
				return x, y
			}
		}
	}

	exclusionCenter = [2]int{x0, y0}
	for i := 0; i < count; i++ {
		x, y := findRandomPoint()
		attachParticle(x, y)
	}
	return &bombs
}
func GetCellNumbers(board *[][]int, bombs *[][2]int) *[][]int {
	rows := len(*board)
	cols := len((*board)[0])

	for bi := range *bombs {
		x0, y0 := (*bombs)[bi][0], (*bombs)[bi][1]
		(*board)[y0][x0] = -1

		// add one score to 8 cells around bombs
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				x, y := x0+i, y0+j
				if (x < cols && x >= 0) && (y < rows && y >= 0) {
					// prevent change the score of bombs cell
					if (*board)[y][x] == -1 {
						continue
					}
					(*board)[y][x]++
				}
			}
		}
	}

	return board
}
