package Tests

import (
	"hello/codeagha_minesweper"
	"testing"
)

func TestAllNeighbors(t *testing.T) {

	cases := []struct {
		in   codeagha_minesweper.NeatCouple
		want map[codeagha_minesweper.NeatCouple]struct{}
	}{

		{[2]int{0, 4}, map[codeagha_minesweper.NeatCouple]struct{}{}},
	}

	cases[0].want[[2]int{0, 3}] = struct{}{}
	cases[0].want[[2]int{1, 3}] = struct{}{}
	cases[0].want[[2]int{0, 4}] = struct{}{}
	cases[0].want[[2]int{0, 4}] = struct{}{}

	for _, c := range cases {

		got := c.in.AllNeighbors([2]int{5, 5})
		for i := range got {
			if _, ok := c.want[i]; ok {

				t.Errorf("err, got : %v | want : %v", got[i], c.want[i])

			}
		}

	}

}
