package entities

type Game struct {
	Board             Board
	Players           []Player
	TurnCounter       int
	CurrentPlayerTurn string
	CardStack         []Card
}

type Board struct {
	Rows [7]Row
}

type Row struct {
	Tiles [7]Tile
}

type Tile struct {
	Piece    int
	PlayerID string
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
