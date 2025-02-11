package minesweeperlib

import (
	"math"
	"math/rand"
	"slices"
)

type Point [2]int8
type Points [][2]int8
type Boardframe [][]int8

func AsPoint(x, y int8) [2]int8 {
	return [2]int8{x, y}
}
func (p *Point) GetComponents() (x, y int8) {
	return p[0], p[1]
}
func (b *Boardframe) GetSize() (cols, rows int8) {
	if len(*b) < 1 {
		return 0, 0
	}
	return int8(len((*b)[0])), int8(len(*b))
}

func GetBoard(pos0 Point) *Boardframe {
	cols, rows := pos0.GetComponents()
	board := make(Boardframe, rows)

	// init the board
	for i := range board {
		board[i] = make([]int8, cols)
	}

	return &board
}
func GetRandomBombs(board *Boardframe, point0 Point, count uint16) *Points {

	bombs := make(Points, 0)
	cols, rows := board.GetSize()

	// The radius of the circle that determines how many points around the point (x0, y0) are forbidden points.
	const radiusFactor = 0.13    // TODO:Add the degree of hardness: the harder, the smaller the radius
	var exclusionCenter = point0 // the point that user clicked on it for first (x0 , y0)

	// The radius of the circle must be specified based on the width of the page.
	calculateExclusionRadius := func() float32 {
		return radiusFactor * float32(cols)
	}

	isInExclusionZone := func(x, y int8) bool {
		radius := calculateExclusionRadius()
		distance := float32(math.Sqrt(math.Pow(float64(x-exclusionCenter[0]), 2) + math.Pow(float64(y-exclusionCenter[1]), 2)))
		return distance < radius
	}

	isFree := func(x, y int8) bool {
		return x >= 0 && y >= 0 && x < cols && y < rows && !slices.Contains(bombs, AsPoint(x, y)) && !isInExclusionZone(x, y)
	}

	attachParticle := func(x, y int8) {
		bombs = append(bombs, AsPoint(x, y))
	}

	findRandomPoint := func() (int8, int8) {
		for {
			x, y := int8(rand.Intn(int(cols))), int8(rand.Intn(int(rows)))
			if isFree(x, y) {
				return x, y
			}
		}
	}

	exclusionCenter = point0
	var i uint16
	for i = 0; i < count; i++ {
		x, y := findRandomPoint()
		attachParticle(x, y)
	}
	return &bombs
}

func GetCellNumbers(board *Boardframe, bombs *Points) *Boardframe {
	cols, rows := board.GetSize()

	for bi := range *bombs {
		x0, y0 := (*bombs)[bi][0], (*bombs)[bi][1]
		(*board)[y0][x0] = -1

		// add one score to 8 cells around bombs
		var i, j int8
		for i = -1; i <= 1; i++ {
			for j = -1; j <= 1; j++ {
				x, y := x0+i, y0+j
				if (x < cols && x >= 0) && (y < rows && y >= 0) {
					// prevent change the score of bombs cell
					if (*board)[y][x] == -1 {
						continue
					}
					(*board)[y][x] += 1
				}
			}
		}
	}

	return board
}

func findZeroNeighbors(point0 Point, openeds *map[Point]struct{}, board *Boardframe) {
	directions := [][2]int8{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	x0, y0 := point0.GetComponents()
	cols, rows := board.GetSize()
	for _, direction := range directions {
		x, y := x0+direction[0], y0+direction[1]
		if x < 0 || x >= cols || y < 0 || y >= rows {
			continue
		}
		thisPos := AsPoint(x, y)
		if _, k := (*openeds)[thisPos]; !k && (*board)[y][x] != -1 {
			(*openeds)[thisPos] = struct{}{}
			if (*board)[y][x] == 0 {
				findZeroNeighbors(thisPos, openeds, board)
			}
		}
	}
}
func GetOpeneds(board *Boardframe, selected Point) Points {
	openeds := make(map[Point]struct{})
	openeds[selected] = struct{}{}
	res := make(Points, 0)
	x0, y0 := selected.GetComponents()
	if (*board)[y0][x0] != 0 {
		for key := range openeds {
			res = append(res, key)
		}
		return res
	}
	findZeroNeighbors(selected, &openeds, board)

	for key := range openeds {
		res = append(res, key)
	}
	return res
}

// States 0: still playing, 1: Winner!, 2: Loser
func GetState(board *Boardframe, bombsCount uint16, flaggeds map[[2]int]bool, point [2]int) int8 {
	if flaggeds == nil {
		if (*board)[point[1]][point[0]] == -1 {
			return 2 // The Game is over !
		}
	}

	if uint16(len(flaggeds)) == bombsCount {
		trueFlags := true
		for flagged := range flaggeds {
			if (*board)[flagged[1]][flagged[0]] != -1 {
				trueFlags = false
				break
			}
		}
		if trueFlags {
			return 1 // Winner
		}
	}

	return 0
}
