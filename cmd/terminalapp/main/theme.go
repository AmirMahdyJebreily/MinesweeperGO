package main

import "fmt"

const unopend = "∙"
const zero = "■"
const bomb = "δ"
const flag = "F"

type Theme struct {
	UsingEscapeCode bool
}

func (t *Theme) ColorsieSymbol(symbol string, escapeCode string, rest string) string {
	if t.UsingEscapeCode {
		return fmt.Sprintf("%v%v%v", escapeCode, symbol, rest)
	}
	return symbol
}

// zero, flag, bomb, unopened default colors
func (t *Theme) DefaultSymbol(symbol string) string {
	switch symbol {
	case zero:
		return t.ColorsieSymbol(zero, "\u001B[38;5;248m", "\u001B[1;0m")
	case flag:
		return t.ColorsieSymbol(flag, "\u001b[102m\u001B[36m", "\u001B[0m")
	case bomb:
		return t.ColorsieSymbol(bomb, "\u001B[101m\u001B[31m", "\u001B[0m")
	case unopend:
		return t.ColorsieSymbol(unopend, "\u001B[38;5;242m", "\u001B[1;0m")
	default:
		return symbol
	}
}

func (t *Theme) ColoriseNumber(n int8) string {
	if n == 0 {
		return t.DefaultSymbol(zero)
	}
	if t.UsingEscapeCode {
		return fmt.Sprintf("\x1b[1;3%vm%v\x1b[0m", n, n)
	}

	return fmt.Sprintf("%v", n)
}
