package minesweeperlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCellNumbers(t *testing.T) {
	type CellNumbersTestCase struct {
		rows, cols int
		bombs      map[int]int
		want       [][]int
	}
	cases := []CellNumbersTestCase{
		{3, 3, map[int]int{1: 1}, [][]int{{1, 1, 1}, {1, -1, 1}, {1, 1, 1}}},
		{3, 3, map[int]int{1: 2}, [][]int{{0, 1, 1}, {0, 1, -1}, {0, 1, 1}}},
		{3, 3, map[int]int{0: 2}, [][]int{{0, 1, -1}, {0, 1, 1}, {0, 0, 0}}},
		{3, 3, map[int]int{0: 2, 1: 1}, [][]int{{1, 2, -1}, {1, -1, 2}, {1, 1, 1}}},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, GetCellNumbers(c.rows, c.cols, c.bombs))
	}

}
