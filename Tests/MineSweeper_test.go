package Tests

import (
	"testing"

	core "github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
)

/*
Test board in case :
______________________________________________
       1   2   3   4   5   6   7   8   9   10|
.......-...-...-...-...-...-...-...-...-...--|
 1 .   0   0   0   0   1   X   1   0   1   X |
 2 .   1   1   0   1   2   2   1   0   2   2 |
 3 .   X   1   0   1   X   1   0   1   2   X |
 4 .   1   1   0   1   1   2   1   2   X   2 |
 5 .   0   0   0   0   0   1   X   2   1   1 |
 6 .   0   0   0   1   1   2   1   2   1   1 |
 7 .   1   1   1   1   X   1   1   2   X   1 |
 8 .   1   X   1   1   1   1   1   X   2   1 |
 ____________________________________________|

*/

var board = core.Init(core.Cpl(8, 10), core.Cpl(5, 6), 11, []core.TCpl{
	{0, 9}, {2, 0},
	{7, 7}, {2, 9},
	{2, 4}, {3, 8},
	{0, 5}, {6, 4},
	{4, 6}, {7, 1},
	{6, 8}})

func TestPointNumber(t *testing.T) {
	cases := []struct {
		in   *core.MineSweeper
		want int
	}{ // arr elem's
		{in: board, want: 2}, // first case
	}

	for _, c := range cases {
		got := c.in.NumberOfPoint(core.Cpl(3, 9))
		if got != c.want {
			t.Errorf("err, got : %v | want : %v", got, c.want)
		}
	}
}

func TestStartPointSafe(t *testing.T) {
	// in this case the c.want should be 8
	// and number of bombs should be maximum value which is this : "(size.I × size.J) - 1"

	cases := []struct {
		in struct {
			sizeI, sizeJ int
			start        core.TCpl
		}
		want int
	}{
		{struct {
			sizeI, sizeJ int
			start        core.TCpl
		}{
			sizeI: 3, sizeJ: 3, start: core.Cpl(1, 1),
		}, 8}, // first case
	}

	for _, c := range cases {
		for i := 0; i < 100; i++ {
			randBoard := core.InitRand(core.Cpl(c.in.sizeI, c.in.sizeJ), c.in.start, (c.in.sizeI*c.in.sizeJ)-1)
			got := randBoard.NumberOfPoint(c.in.start)
			if got != c.want {
				t.Errorf("err, got : %v | want : %v", got, c.want)
				break
			}
		}
	}
}
