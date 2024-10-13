package Tests

import (
	"testing"

	"github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
)

var board = GameCore.MineSweeper{
	BombsCount: 11,
	Size:       GameCore.Cpl(8, 10),
	Bombs: []GameCore.TCpl{
		{0, 9}, {2, 0},
		{7, 7}, {2, 9},
		{2, 4}, {3, 8},
		{0, 5}, {6, 4},
		{4, 6}, {7, 1},
		{6, 8}},
}

func TestPointNumber(t *testing.T) {
	cases := []struct {
		in   GameCore.MineSweeper
		want int
	}{ // arr elem's
		{in: board, want: 2}, // index 0
	}

	for _, c := range cases {
		got := c.in.NumberOfPoint(GameCore.Cpl(3, 9))
		if got != c.want {
			t.Errorf("err, got : %v | want : %v", got, c.want)
		}
	}
}
