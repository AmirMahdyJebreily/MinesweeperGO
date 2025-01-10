package main

import (
	"fmt"
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
)

const untaped = "∙"
const bomb = "\u001b[41m\u001B[35mX\u001B[0m"

func Sprintgridf(board *[][]int, bombs *[][2]int) string {
	cols, rows := len(*board), len((*board)[0])
	res := ""
	for i := 0; i < cols; i++ {
		if i == 0 {
			res += "      "
			for j := 0; j < rows; j++ {
				res += fmt.Sprintf(" %02d  ", j+1)
			}
			res += fmt.Sprintf("\n\n")

		}
		res += fmt.Sprintf(" %02d     ", i+1)
	lines:
		for j := 0; j < rows; j++ {
			if bombs != nil {
				thisPos := [2]int{i, j}
				for _, bom := range *bombs {
					if thisPos == bom {
						res += fmt.Sprintf("%v    ", bomb)
						continue lines
					}
				}
			}
			res += fmt.Sprintf("%v    ", (*board)[i][j])

		}
		res += "\n"
	}
	return res
}

func main() {
	var cols, rows int
	fmt.Println("Wellcome to CodeAgha's MineSweeper Game in terminal")
	for {
		fmt.Print("Enter The Columns,Rows: ")
		_, scanError := fmt.Scanf("%v,%v", &cols, &rows)
		if scanError != nil {
			fmt.Println("Please Input values in format above: Columns,Rows")
			continue
		}
		break
	}
	var bombsCount int
	for {
		fmt.Print("Enter The count of bombs: ")
		_, scanError := fmt.Scanf("%v", &bombsCount)
		if scanError != nil {
			fmt.Println("Please Input values in format above: Columns,Rows")
			continue
		}
		break
	}

	board := minesweeperlib.GetBoard(cols, rows)
	fmt.Println(Sprintgridf(board, nil))
	var x0, y0 int
	for {
		fmt.Print("Select: ")
		_, scanError := fmt.Scanf("%v,%v", &x0, &y0)
		if scanError != nil {
			fmt.Println("Please Input values in format above: Columns,Rows")
			continue
		}
		break
	}
	bombs := minesweeperlib.GetRandomBombs(cols, rows, x0, y0, bombsCount)
	board = minesweeperlib.GetCellNumbers(board, bombs)
	fmt.Println(Sprintgridf(board, bombs))

	//for {
	//
	//}
}
