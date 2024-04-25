package entities

type Game struct {
	Board Board
}

type Board struct {
	Rows [7]Row
}

type Row struct {
	Tiles [7]Tile
}

type Tile struct {
	Piece int
}

type TileCoords struct {
	Row  int
	Tile int
}

type Player struct {
	Code           string
	AlivePieces    []TileCoords
	DefeatedPieces []int
	DrawnCard      Card
}

type Card struct {
	Code  int
	Value int
}
