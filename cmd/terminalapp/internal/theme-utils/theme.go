package theme

import "fmt"

const (
	Unopend = "∙"
	Zero    = "■"
	Bomb    = "δ"
	Flag    = "F"
)

var UsingEscapeCode bool

func ColorsieSymbol(symbol string, escapeCode string, rest string) string {
	if UsingEscapeCode {
		return fmt.Sprintf("%v%v%v", escapeCode, symbol, rest)
	}
	return symbol
}

// zero, flag, bomb, unopened default colors
func DefaultSymbol(symbol string) string {
	switch symbol {
	case Zero:
		return ColorsieSymbol(Zero, "\u001B[38;5;248m", "\u001B[1;0m")
	case Flag:
		return ColorsieSymbol(Flag, "\u001b[102m\u001B[36m", "\u001B[0m")
	case Bomb:
		return ColorsieSymbol(Bomb, "\u001B[101m\u001B[31m", "\u001B[0m")
	case Unopend:
		return ColorsieSymbol(Unopend, "\u001B[38;5;242m", "\u001B[1;0m")
	default:
		return symbol
	}
}

func ColoriseNumber(n int8) string {
	if n == 0 {
		return DefaultSymbol(Zero)
	}
	if UsingEscapeCode {
		return fmt.Sprintf("\x1b[1;3%vm%v\x1b[0m", n, n)
	}

	return fmt.Sprintf("%v", n)
}
