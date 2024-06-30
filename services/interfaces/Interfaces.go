package interfaces

import (
	"luckyChess/entities"
)

type IGameStoreService interface {
	NewGame(startingSet entities.BoardTemplate) entities.Game
	GetGame(code string) (entities.Game, error)
	KillGame(code string) error
	ApplyChanges(code string, game entities.Game) error
}

type IGameStateService interface {
	MovePiece(game *entities.Game, playerCode string, pieceCoords entities.TileCoords, targetCoords entities.TileCoords) error
	GetMoveset(game *entities.Game, playerCode string, pieceCoords entities.TileCoords) ([]entities.TileCoords, error)
	DrawCard(game *entities.Game, playerCode string) (*entities.Card, error)
	CheckWin(game *entities.Game) (string, error) /*return player code if a player has won, otherwise returns empty string*/
}

type IGameTemplatesService interface {
	GetTemplate(name string) entities.BoardTemplate
	NewTemplate(name string, template entities.BoardTemplate) error
	AtlerTemplate(name string, template entities.BoardTemplate) error
	RemoveTemplate(name string, template entities.BoardTemplate) error
}
