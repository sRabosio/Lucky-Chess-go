package interfaces

import(
	"luckyChess/entities"
)

type IGameStoreService interface {
	NewGame() Game
	GetGame(code string) Game
	KillGame(code string) bool
	ApplyChanges(code string, game Game) bool
}

type IGameStateService interface{
	MovePiece(game &Game, pieceCoords TileCoords, targetCoords TileCoords) bool
	GetMoveset(game &Game, playerCode string, pieceCoords TileCoords) TileCoords[]
	DrawCard(game &Game, playerCode string) &Card
	CheckWin(game &Game) string /*return player code if a player has won, otherwise returns empty string*/
}
