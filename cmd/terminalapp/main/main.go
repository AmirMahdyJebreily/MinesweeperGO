package main

import (
	"fmt"
	"slices"
	"strings"

	mnsw "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/eiannone/keyboard"
)

var theme Theme

func SprintCell(data string, selected bool) string {
	if selected {
		if theme.UsingEscapeCode {
			return fmt.Sprintf("\u001B[5m[\u001B[25m%v\u001B[5m]\x1b[25m", data)
		}

		return fmt.Sprintf("[%v]", data)
	}

	if theme.UsingEscapeCode {
		if data == flag {
			return fmt.Sprintf("\u001B[102m\u001B[36m[%v]\u001B[0m", flag)
		}
		if data == bomb {
			return fmt.Sprintf("\u001B[101m\u001B[31m[%v]\u001B[0m", bomb)
		}
	}

	return fmt.Sprintf(" %v ", data)
}

func Sprintgridf(board *mnsw.Boardframe, bombsCount int8, flagged *map[mnsw.Point]bool, oppend *mnsw.Points, selected [2]int8, messages string) *strings.Builder {
	rows, cols := board.GetSize()
	var res strings.Builder
	res.WriteString("\033[H\033[J") // Clear the old screen before print new board
	res.WriteString(fmt.Sprintf(" [Size: %v×%v] [Bombs: %v] [Flags: %v]\n", cols, rows, bombsCount, bombsCount-int8(len(*flagged))))
	var i, j int8
	for i = rows - 1; i >= 0; i-- {
	lines:
		for j = 0; j < cols; j++ {
			isSelected := [2]int8{j, i} == selected

			if _, isflag := (*flagged)[[2]int8{j, i}]; isflag {
				res.WriteString(SprintCell(theme.DefaultSymbol(flag), isSelected))
				continue lines
			}

			if (oppend) != nil {
				if !slices.Contains(*oppend, [2]int8{j, i}) {
					res.WriteString(SprintCell(theme.DefaultSymbol(unopend), isSelected))
					continue lines
				}
			}

			if (*board)[i][j] == -1 {
				res.WriteString(SprintCell(theme.DefaultSymbol(bomb), isSelected))
				continue lines
			}

			res.WriteString(SprintCell(theme.ColoriseNumber((*board)[i][j]), isSelected))
		}
		res.WriteString("\n")
	}
	res.WriteString(fmt.Sprintf("\n%v", messages))

	return &res
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Print("Do you want to use ANSI Escape codes [(y)Yes/(n)No] ('y' is default) ? ")
		input := "y"
		_, scanError := fmt.Scanf("%v\n", &input)
		if scanError != nil || strings.Trim(input, " ") != "n" {
			theme.UsingEscapeCode = true
			break
		} else if strings.Trim(input, " ") == "n" {
			theme.UsingEscapeCode = false
			break
		}
		theme.UsingEscapeCode = true
		break
	}

	var cols, rows int8
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
	var bombsCount int8
	for {
		suggested := int8(int(float64(cols)*float64(rows)*0.21) - int(1))
		fmt.Printf("\nEnter The count of bombs (default %v bombs): ", suggested)
		_, scanError := fmt.Scanf("%v\n", &bombsCount)
		if scanError != nil {
			bombsCount = suggested
			break
		}
		break
	}
	fmt.Println("\n\u001B[?1049h\u001B[H\u001B[J\u001B[?25l") // save screen
	board := mnsw.GetBoard(mnsw.AsPoint(cols, rows))
	flaggeds := make(map[mnsw.Point]bool, bombsCount)
	selected := mnsw.AsPoint(cols/2, rows/2)
	message := "[Arrows: Move] [O & Enter: Open Cell]\n[F & Space: Flag] [Q & ESC: Quit]\nSelect a cell to start game"
	fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, message)).String())
	var x0, y0 int8
	var bombs *mnsw.Points = nil
	var oppend mnsw.Points = nil
	inGame := true
	for inGame {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowLeft && selected[0] > 0 {
			selected[0]--
		}

		if key == keyboard.KeyArrowRight && selected[0] < cols-1 {
			selected[0]++
		}

		if key == keyboard.KeyArrowUp && selected[1] < rows-1 {
			selected[1]++
		}

		if key == keyboard.KeyArrowDown && selected[1] > 0 {
			selected[1]--
		}

		if char == 'q' || char == 'Q' || key == keyboard.KeyEsc {
			fmt.Print("\u001B[?1049l\u001B[?25h")
			break
		}

		if char == 'f' || char == 'F' || key == keyboard.KeySpace {
			if val, isflag := flaggeds[selected]; isflag || val {
				delete(flaggeds, selected)
			} else if int8(len(flaggeds)) < bombsCount {
				flaggeds[selected] = true

				state := mnsw.GetState(board, bombsCount, flaggeds, selected)
				if state == 1 {
					message = "You Win :)"
					inGame = false
					fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, message)).String())
					fmt.Println("Press something to exit\n")
					fmt.Scanln()
					fmt.Print("\u001B[?1049l\u001B[?25h")
					break
				}
			}
		}
		if char == 'o' || char == 'O' || key == keyboard.KeyEnter {
			x0, y0 = selected[0], selected[1]
			if bombs == nil {
				oppend = make([][2]int8, 0)
				bombs = mnsw.GetRandomBombs(board, mnsw.AsPoint(x0, y0), bombsCount)
				board = mnsw.GetCellNumbers(board, bombs)
				message = "[Arrows: Move] [O & Enter: Open Cell]\n[F & Space: Flag] [Q & ESC: Quit]"
			}
			oppend = slices.Concat(oppend, mnsw.GetOpeneds(board, mnsw.AsPoint(x0, y0)))

			state := mnsw.GetState(board, bombsCount, nil, selected)
			if state == 2 {
				message = "Game Over :("
				inGame = false
				fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, message)).String())
				fmt.Println("Press something to exit\n")
				fmt.Scanln()
				fmt.Print("\u001B[?1049l\u001B[?25h")
				break
			}
		}

		// update screen
		fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, &oppend, selected, message)).String())
		fmt.Print()
	}
}
