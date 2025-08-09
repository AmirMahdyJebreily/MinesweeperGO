package internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AmirMahdyJebreily/MinesweeperGO/cmd/terminalapp/internal/theme-utils"
	mnsw "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
)

func sprintCell(data string, selected bool) string {
	if selected {
		if theme.UsingEscapeCode {
			return fmt.Sprintf("\u001B[5m[\u001B[25m%v\u001B[5m]\x1b[25m", data)
		}

		return fmt.Sprintf("[%v]", data)
	}

	if theme.UsingEscapeCode {
		if data == theme.Flag {
			return fmt.Sprintf("\u001B[102m\u001B[36m[%v]\u001B[0m", theme.Flag)
		}
		if data == theme.Bomb {
			return fmt.Sprintf("\u001B[101m\u001B[31m[%v]\u001B[0m", theme.Bomb)
		}
	}

	return fmt.Sprintf(" %v ", data)
}

func PrintBoard(board *mnsw.Boardframe, bombsCount int8, flagged *mnsw.Points, oppend *mnsw.Points, selected [2]int8, messages string) *strings.Builder {
	rows, cols := board.GetSize()
	var res strings.Builder
	res.WriteString("\033[H\033[J") // Clear the old screen before print new board
	res.WriteString(fmt.Sprintf(" [Size: %v×%v] [Bombs: %v] [Flags: %v | %v]\n", cols, rows, bombsCount, int8(len(*flagged)), bombsCount-int8(len(*flagged))))
	var i, j int8
	for i = rows - 1; i >= 0; i-- {
	lines:
		for j = 0; j < cols; j++ {
			isSelected := [2]int8{j, i} == selected

			if slices.Contains(*flagged, [2]int8{j, i}) {
				res.WriteString(sprintCell(theme.DefaultSymbol(theme.Flag), isSelected))
				continue lines
			}

			if (oppend) != nil {
				if !slices.Contains(*oppend, [2]int8{j, i}) {
					res.WriteString(sprintCell(theme.DefaultSymbol(theme.Unopend), isSelected))
					continue lines
				}
			}

			if (*board)[i][j] == -1 {
				res.WriteString(sprintCell(theme.DefaultSymbol(theme.Bomb), isSelected))
				continue lines
			}

			res.WriteString(sprintCell(theme.ColoriseNumber((*board)[i][j]), isSelected))
		}
		res.WriteString("\n")
	}
	res.WriteString(fmt.Sprintf("\n%v", messages))

	return &res
}
