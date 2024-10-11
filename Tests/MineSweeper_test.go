package Tests

import (
	"hello/codeagha_minesweper/GameCore"
	"testing"
)

func TestPointNumber(t *testing.T) {
	cases := []struct {
		in   GameCore.MineSweeper
		want int
	}{

		{GameCore.MineSweeper{
			Size:       [2]int{3, 3},
			BombsCount: 3,
			Bombs: []GameCore.NeatCouple{
				{0, 0},
				{1, 1},
				{1, 2},
			},
		}, 3},
	}

	for _, c := range cases {
		got := c.in.NumberOfPoint([2]int{0, 1})
		if got != c.want {
			t.Errorf("err, got : %v | want : %v", got, c.want)
		}
	}
}
