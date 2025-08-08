package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AmirMahdyJebreily/MinesweeperGO/cmd/terminalapp/internal"
	mnsw "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/eiannone/keyboard"
)

var theme Theme

func main() {

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
	fmt.Println((*internal.PrintBoard(board, bombsCount, &flaggeds, nil, selected, message)).String())
	var x0, y0 int8
	var bombs *mnsw.Points = nil
	var oppend mnsw.Points = nil
	inGame := true
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
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
				break
			}

			if int8(len(flaggeds)) < bombsCount {
				flaggeds[selected] = true

				state := mnsw.GetState(board, bombsCount, flaggeds, selected)
				if state == 1 {
					message = "You Win :)"
					inGame = false
					err := keyboard.Close()
					if err != nil {
						return
					}
					fmt.Println((*internal.PrintBoard(board, bombsCount, &flaggeds, nil, selected, message)).String())
					fmt.Println("Press something to exit\n")
					fmt.Scanln()
					fmt.Print("\u001B[?1049l\u001B[?25h")
					break
				}
				break
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
				fmt.Println((*internal.PrintBoard(board, bombsCount, &flaggeds, nil, selected, message)).String())
				fmt.Println("Press something to exit\n")
				fmt.Scanln()
				fmt.Print("\u001B[?1049l\u001B[?25h")
				break
			}
		}

		// update screen
		fmt.Println((*internal.PrintBoard(board, bombsCount, &flaggeds, &oppend, selected, message)).String())
		fmt.Print()
	}
}
