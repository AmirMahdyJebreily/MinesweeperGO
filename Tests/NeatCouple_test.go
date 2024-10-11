package Tests

import (
	"testing"

	"github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
)

func TestAllNeighbors(t *testing.T) {

	cases := []struct {
		in   GameCore.NeatCouple
		want map[GameCore.NeatCouple]struct{}
	}{

		{[2]int{0, 4}, map[GameCore.NeatCouple]struct{}{}},
	}

	cases[0].want[[2]int{0, 3}] = struct{}{}
	cases[0].want[[2]int{1, 3}] = struct{}{}
	cases[0].want[[2]int{1, 4}] = struct{}{}
	cases[0].want[[2]int{0, 4}] = struct{}{}

	for _, c := range cases {

		got := c.in.AllNeighbors([2]int{5, 5})
		for i := range got {
			if _, ok := c.want[i]; !ok {

				t.Errorf("err, got : %v | want : %v", i, c.want[i])

			}
		}

	}

}
