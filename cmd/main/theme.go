package main

import "fmt"

const unoppend = "\u001B[38;5;242m∙\u001B[1;0m"
const zero = "\u001B[38;5;248m■\u001B[1;0m"
const bomb = "\u001b[101m\u001B[31mX\u001B[0m"
const flag = "\u001b[102m\u001B[36mF\u001B[0m"

type Theme struct {
	UsingEscapeCode bool
}

func (t *Theme) ColorsieSymbol(symbol string, escapeCode string, rest string) string {
	if t.UsingEscapeCode {
		return fmt.Sprintf("%v%v%%v", escapeCode, symbol, rest)
	}
	return symbol
}

func (t *Theme) ColoriseNumber(n int) string {
	if n == 0 {
		return zero
	}
	if t.UsingEscapeCode {
		return fmt.Sprintf("\x1b[1;3%vm%v\x1b[0m", n, n)
	}

	return fmt.Sprintf("%v", n)
}
