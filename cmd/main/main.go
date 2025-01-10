package main

import (
	"fmt"
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"strings"
)

const untaped = "∙"
const bomb = "\u001b[41m\u001B[35mX\u001B[0m"

// use \x1b[1;31m ansi escape code for colorise
func coloriseNumber(n int) string {
	return fmt.Sprintf("\x1b[1;3%vm%v\x1b[1;0m", n, n)
}

func Sprintgridf(board *[][]int, bombs *[][2]int) *strings.Builder {
	rows, cols := len(*board), len((*board)[0])
	var res strings.Builder
	res.WriteString("\033[H\033[2J") // Clear the old screen before print new board
	if bombs != nil {
		res.WriteString(fmt.Sprintf(" .:: [Size: %v×%v] [Bombs: %v]\n", cols, rows, len(*bombs)))
	}
	for i := rows - 1; i >= 0; i-- {
		res.WriteString(fmt.Sprintf(" %2d     ", i+1))
	lines:
		for j := 0; j < cols; j++ {
			if (*board)[i][j] == -1 {
				res.WriteString(fmt.Sprintf("%v    ", bomb))
				continue lines
			}

			res.WriteString(fmt.Sprintf("%v    ", coloriseNumber((*board)[i][j])))

		}
		res.WriteString("\n")
	}
	res.WriteString(fmt.Sprintf("\n      "))
	for j := 0; j < cols; j++ {
		res.WriteString(fmt.Sprintf(" %2d  ", j+1))
	}
	res.WriteString(fmt.Sprintf("\n"))

	return &res
}

func main() {
	var cols, rows int
	fmt.Println("Wellcome to CodeAgha's MineSweeper Game in terminal")
	for {
		fmt.Print("Enter The Columns,Rows: ")
		_, scanError := fmt.Scanf("%v,%v\n", &cols, &rows)
		if scanError != nil {
			fmt.Println("Please Input values in format: Columns,Rows")
			continue
		}
		break
	}
	var bombsCount int
	for {
		fmt.Print("Enter The count of bombs: ")
		_, scanError := fmt.Scanln(&bombsCount)
		if scanError != nil {
			fmt.Println("Please Input a number")
			continue
		}
		break
	}

	board := minesweeperlib.GetBoard(cols, rows)
	fmt.Println((*Sprintgridf(board, nil)).String())
	var x0, y0 int
	for {
		fmt.Print("Select a cell to start game: ")
		_, scanError := fmt.Scanf("%v,%v", &x0, &y0)
		if scanError != nil {
			fmt.Println("Please Input values in format above: Columns,Rows")
			continue
		}
		break
	}
	bombs := minesweeperlib.GetRandomBombs(cols, rows, x0-1, y0-1, bombsCount)
	board = minesweeperlib.GetCellNumbers(board, bombs)
	fmt.Println((*Sprintgridf(board, bombs)).String())

	//for {
	//
	//}
}
