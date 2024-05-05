package interfaces

import (
	"luckyChess/entities"
)

type IGameStoreService interface {
	NewGame(startingSet entities.Board) entities.Game
	GetGame(code string) entities.Game
	KillGame(code string) bool
	ApplyChanges(code string, game entities.Game) bool
}

type IGameStateService interface {
	MovePiece(game *entities.Game, pieceCoords entities.TileCoords, targetCoords entities.TileCoords) bool
	GetMoveset(game *entities.Game, playerCode string, pieceCoords entities.TileCoords)
	DrawCard(game *entities.Game, playerCode string) *entities.Card
	CheckWin(game *entities.Game) string /*return player code if a player has won, otherwise returns empty string*/
}
