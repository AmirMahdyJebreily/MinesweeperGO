package minesweeperlib_test

import (
	mnsw "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestGetBoard(t *testing.T) {
	cases := [3](mnsw.Point){
		{3, 3},
		{3, 6},
		{6, 3},
	}
	for _, want := range cases {
		res := *mnsw.GetBoard(want)
		var actual mnsw.Point = mnsw.AsPoint(int8(len((res)[0])), int8(len(res)))
		assert.Equal(t, want, actual)
	}
}

func TestGetRandomBombs(t *testing.T) {
	board := mnsw.GetBoard(mnsw.AsPoint(3, 3))
	bombs := *mnsw.GetRandomBombs(board, mnsw.AsPoint(1, 1), 8)
	assert.False(t, slices.Contains(bombs, mnsw.AsPoint(1, 1)))
}

func TestGetCellNumbers(t *testing.T) {
	type CellNumbersTestCase struct {
		rows, cols int8
		bombs      mnsw.Points
		want       mnsw.Boardframe
	}
	cases := [4]CellNumbersTestCase{
		{3, 3, [][2]int8{{1, 1}}, [][]int8{{1, 1, 1}, {1, -1, 1}, {1, 1, 1}}},
		{3, 3, [][2]int8{{2, 1}}, [][]int8{{0, 1, 1}, {0, 1, -1}, {0, 1, 1}}},
		{3, 3, [][2]int8{{2, 0}}, [][]int8{{0, 1, -1}, {0, 1, 1}, {0, 0, 0}}},
		{3, 3, [][2]int8{{2, 0}, {1, 1}}, [][]int8{{1, 2, -1}, {1, -1, 2}, {1, 1, 1}}},
	}
	for _, c := range cases {
		assert.Equal(t, &c.want, mnsw.GetCellNumbers(mnsw.GetBoard(mnsw.AsPoint(c.cols, c.rows)), &c.bombs))
	}

}

func TestGetOpeneds(t *testing.T) {
	o := mnsw.GetOpeneds(mnsw.GetBoard(mnsw.AsPoint(3, 3)), mnsw.AsPoint(2, 2))
	assert.Contains(t, o, [2]int8{2, 2})
}
