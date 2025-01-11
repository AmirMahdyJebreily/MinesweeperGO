package minesweeperlib_test

import (
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestGetBoard(t *testing.T) {
	cases := [3]([2]int){
		{3, 3},
		{3, 6},
		{6, 3},
	}
	for _, want := range cases {
		res := *minesweeperlib.GetBoard(want[0], want[1])
		actual := [2]int{len((res)[0]), len(res)}
		assert.Equal(t, want, actual)
	}
}

func TestGetRandomBombs(t *testing.T) {
	bombs := *minesweeperlib.GetRandomBombs(3, 3, 1, 1, 8)
	assert.False(t, slices.Contains(bombs, [2]int{1, 1}))
}

func TestGetCellNumbers(t *testing.T) {
	type CellNumbersTestCase struct {
		rows, cols int
		bombs      [][2]int
		want       [][]int
	}
	cases := [4]CellNumbersTestCase{
		{3, 3, [][2]int{{1, 1}}, [][]int{{1, 1, 1}, {1, -1, 1}, {1, 1, 1}}},
		{3, 3, [][2]int{{2, 1}}, [][]int{{0, 1, 1}, {0, 1, -1}, {0, 1, 1}}},
		{3, 3, [][2]int{{2, 0}}, [][]int{{0, 1, -1}, {0, 1, 1}, {0, 0, 0}}},
		{3, 3, [][2]int{{2, 0}, {1, 1}}, [][]int{{1, 2, -1}, {1, -1, 2}, {1, 1, 1}}},
	}
	for _, c := range cases {
		assert.Equal(t, &c.want, minesweeperlib.GetCellNumbers(minesweeperlib.GetBoard(c.cols, c.rows), &c.bombs))
	}

}

func TestGetOpeneds(t *testing.T) {
	point := [2]int{2, 2}
	o := minesweeperlib.GetOpeneds(minesweeperlib.GetBoard(3, 3), point)
	assert.Contains(t, o, point)
}
