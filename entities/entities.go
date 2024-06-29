package entities

import eChess "luckyChess/entities/EChess"

type Game struct {
	Board             Board
	Players           []Player
	TurnCounter       int
	CurrentPlayerTurn string
	CardStack         []Card
}

type BoardViewState struct {
	Game      Game
	SelectedX int
	SelectedY int
}

type Board struct {
	Rows [8]Row
}

type Row struct {
	Tiles [8]Tile
}

type Tile struct {
	Piece    eChess.EChess
	PlayerID string
	State    string
}

type TileCoords struct {
	Row  int
	Tile int
}

type Player struct {
	Code           string
	DefeatedPieces []int
	DrawnCard      Card
}

type Card struct {
	Code  int
	Value int
}

type BoardTemplate struct {
	Template []map[int]map[int]string
}
