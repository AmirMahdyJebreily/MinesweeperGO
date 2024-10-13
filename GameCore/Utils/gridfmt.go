package utils

import (
	"fmt"
	"strings"

	core "github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
)

const untaped = "∙"
const bomb = "\u001b[41m\u001B[35mX\u001B[0m"

func Sprintgridf(m *core.MineSweeper) string {
	var res strings.Builder
	for i := 0; i < m.Size[0]; i++ {
		if i == 0 {
			res.WriteString("      ")
			for j := 0; j < m.Size[1]; j++ {
				res.WriteString(fmt.Sprintf(" %v  ", j+1))
			}
			res.WriteString(fmt.Sprintln("\n"))

		}
		res.WriteString(fmt.Sprintf(" %v     ", i+1))
	lines:
		for j := 0; j < m.Size[1]; j++ {

			thisPos := [2]int{i, j}
			for _, bom := range m.Bombs {
				if thisPos == bom {
					res.WriteString(fmt.Sprintf("%v   ", bomb))
					continue lines
				}
			}
			res.WriteString(fmt.Sprintf("%v   ", m.NumberOfPoint([2]int{i, j})))

		}
		res.WriteString("\n")
	}
	return res.String()
}
