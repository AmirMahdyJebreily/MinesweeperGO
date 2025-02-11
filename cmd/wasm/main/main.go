package main

import (
	"fmt"
	mnsw "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"go/types"
	"slices"
	"syscall/js"
)

//func JS_GetBoard() js.Func {
//
//	return js.FuncOf(func(this js.Value, args []js.Value) any {
//		if len(args) != 2 {
//			return ArgumentError.Error()
//		}
//		cols, rows := args[0].Int(), args[1].Int()
//
//		return js.ValueOf(Convert2dArrayToJsArray(minesweeperlib.GetBoard(cols, rows)))
//	})
//}
//
//func JS_GetRandomBombs() js.Func {
//	var ArgumentError = js.ValueError{
//		Method: "[minesweeperlib_JS_GetRandomBombs]: Input arguments are not supported\nplease see documents again at (https://github.com/AmirMahdyJebreily/MinesweeperGO/tree/main/cmd/wasm/main)",
//		Type:   types.Int,
//	}
//	return js.FuncOf(func(this js.Value, args []js.Value) any {
//		if len(args) != 5 {
//			return ArgumentError.Error()
//		}
//		cols, rows, x0, y0, count := args[0].Int(), args[1].Int(), args[2].Int(), args[3].Int(), args[4].Int()
//		return js.ValueOf(Convert2dPosArrayToJsArray(minesweeperlib.GetRandomBombs(cols, rows, x0, y0, count)))
//	})
//}
//
//func JS_GetCellNumbers() js.Func {
//	var ArgumentError = js.ValueError{
//		Method: "[minesweeperlib_JS_GetRandomBombs]: Input arguments are not supported\nplease see documents again at (https://github.com/AmirMahdyJebreily/MinesweeperGO/tree/main/cmd/wasm/main)",
//		Type:   types.Int,
//	}
//	return js.FuncOf(func(this js.Value, args []js.Value) any {
//		if len(args) != 2 {
//			return ArgumentError.Error()
//		}
//
//		board, bombs := args[0].Int(), args[1].Int()
//		return js.ValueOf(Convert2dArrayToJsArray(minesweeperlib.GetCellNumbers(board, bombs)))
//	})
//}

var cols, rows, bombsCount int8
var selcted mnsw.Point
var board *mnsw.Boardframe
var bombs *mnsw.Points = nil
var oppend mnsw.Points = nil

func main() {
	var ArgumentError = func(text string) js.ValueError {
		return js.ValueError{
			Method: fmt.Sprintf("Error: %v\nplease see documents again at (https://github.com/AmirMahdyJebreily/MinesweeperGO/tree/main/cmd/wasm/main)", text),
			Type:   types.Int,
		}
	}
	ch := make(chan interface{}, 0)
	js.Global().Set("startGame", js.FuncOf(
		func(this js.Value, args []js.Value) any {
			if len(args) != 4 {
				jsErr := ArgumentError("Start arguments error")
				return jsErr.Error()
			}
			cols, rows, bombsCount = int8(args[0].Int()), int8(args[1].Int()), int8(args[2].Int())
			selcted = mnsw.AsPoint(int8(args[2].Int()), int8(args[3].Int()))
			board = mnsw.GetBoard(mnsw.AsPoint(cols, rows))
			bombs = mnsw.GetRandomBombs(board, selcted, bombsCount)
			board = mnsw.GetCellNumbers(board, bombs)
			oppend = slices.Concat(oppend, mnsw.GetOpeneds(board, selcted))
			return 0
		}))
	<-ch
}
