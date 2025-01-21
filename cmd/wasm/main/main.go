package main

import (
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"go/types"
	"syscall/js"
)

func JS_GetBoard() js.Func {
	var ArgumentError = js.ValueError{
		Method: "[MineSweeperLib_Js_Getboard]: Input arguments are not supported\nplease see documents again at",
		Type:   types.Int,
	}
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return ArgumentError.Error()
		}
		cols, rows := args[0].Int(), args[1].Int()
		return minesweeperlib.GetBoard(cols, rows)
	})
}

func JS_GetRandomBombs() js.Func {

}

func main() {

}
