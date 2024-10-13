package main

import (
	"fmt"

	"github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
)

const untaped = "∙"
const bomb = "\u001b[41m\u001B[35mX\u001B[0m"

func Sprintgridf(m *GameCore.MineSweeper) string {
	res := ""
	for i := 0; i < m.Size[0]; i++ {
		if i == 0 {
			res += "      "
			for j := 0; j < m.Size[1]; j++ {
				res += fmt.Sprintf(" %v  ", j+1)
			}
			res += fmt.Sprintf("\n\n")

		}
		res += fmt.Sprintf(" %v     ", i+1)
	lines:
		for j := 0; j < m.Size[1]; j++ {

			thisPos := [2]int{i, j}
			for _, bom := range m.Bombs {
				if thisPos == bom {
					res += fmt.Sprintf("%v   ", bomb)
					continue lines
				}
			}
			res += fmt.Sprintf("%v   ", m.NumberOfPoint([2]int{i, j}))

		}
		res += "\n"
	}
	return res
}

func main() {
	g := GameCore.Init(GameCore.Cpl(8, 10), GameCore.Cpl(5, 6), 11)

	fmt.Printf("%v", Sprintgridf(g))
	fmt.Printf("\n%#v", g.Bombs)
}
