package eChess

import "strings"

type EChess int

// enum
const (
	_ EChess = iota
	PAWN
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

var _names = []string{
	"",
	"PAWN",
	"ROOK",
	"KNIGHT",
	"BISHOP",
	"QUEEN",
	"KING",
}

func (num EChess) String() string {
	return _names[num]
}

func Parse(query string) EChess {
	for index, name := range _names {
		if name == strings.ToUpper(query) {
			return EChess(index)
		}
	}
	return 0
}
