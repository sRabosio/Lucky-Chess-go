package eSessionStatus

import "strings"

type ESessionStatus int

// enum
const (
	_ ESessionStatus = iota
	NO_USER
	NO_GAME
	IN_GAME
)

var _names = []string{
	"",
	"NO_USER",
	"NO_GAME",
	"IN_GAME",
}

func (num ESessionStatus) String() string {
	return _names[num]
}

func Parse(query string) ESessionStatus {
	for index, name := range _names {
		if name == strings.ToUpper(query) {
			return ESessionStatus(index)
		}
	}
	return 0
}
