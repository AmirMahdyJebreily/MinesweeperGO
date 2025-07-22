package main

import (
	lib "github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"iter"
)

func Convert2dPosArrayToJsArray[T any](goArr *[][2]T) iter.Seq2[int, []interface{}] {
	return func(yeild func(index int, items []interface{}) bool) {
		for i := range *goArr {
			r := make([]interface{}, len((*goArr)[i]))
			for j := range (*goArr)[i] {
				r[j] = (*goArr)[i][j]
			}
			if !yeild(i, r) {
				return
			}
		}
	}
}

func ConvertBoardToJsArray(b *lib.Boardframe) []interface{} {
	res := make([]interface{}, len((*b)))
	for i := range *b {
		r := make([]interface{}, len((*b)[i]))
		for j := range (*b)[i] {
			r[j] = (*b)[i][j]
		}
		res[i] = r
	}
	return res
}

func Convert2dArrayToJsArray[T any](goArr *[][]T) iter.Seq2[int, []interface{}] {
	return func(yeild func(index int, items []interface{}) bool) {
		for i := range *goArr {
			r := make([]interface{}, len((*goArr)[i]))
			for j := range (*goArr)[i] {
				r[j] = (*goArr)[i][j]
			}
			if !yeild(i, r) {
				return
			}
		}
	}
}
