package gameStateService

import (
	"errors"
	"luckyChess/entities"
	eChess "luckyChess/entities/EChess"
)

type GameStateService struct{}

func New() *GameStateService {
	return &GameStateService{}
}

func (g GameStateService) MovePiece(game *entities.Game, pieceCoords entities.TileCoords, targetCoords entities.TileCoords) bool {
	panic("not implemented")
}

func (g GameStateService) GetMoveset(game *entities.Game, playerCode string, pieceCoords entities.TileCoords) ([]entities.TileCoords, error) {
	x, y := pieceCoords.Tile, pieceCoords.Row

	selectedTile := &game.Board.Rows[y].Tiles[x]

	res := []entities.TileCoords{}
	//TODO CHECK PLAYER TYPE: ex(first or second/ black or white)

	rowLen := len(game.Board.Rows)

	if selectedTile.Piece < 1 {
		return []entities.TileCoords{}, nil
	}

	if playerCode != selectedTile.PlayerID {
		return res, errors.New("invalid player")
	}

	switch selectedPiece := selectedTile.Piece; selectedPiece {
	case eChess.PAWN:

		if y >= rowLen-1 {
			break
		}

		//NB: this is bottom player prospective

		//in front
		if game.Board.Rows[y-1].Tiles[x].Piece < 1 {
			res = append(res, entities.TileCoords{Tile: x, Row: y - 1})
		}

		//right
		if len(game.Board.Rows[y-1].Tiles) < x && game.Board.Rows[y-1].Tiles[x-1].Piece > 0 {
			res = append(res, entities.TileCoords{Tile: x - 1, Row: y})
		}

		//left
		if len(game.Board.Rows[y-1].Tiles) < x && game.Board.Rows[y-1].Tiles[x+1].Piece > 0 {
			res = append(res, entities.TileCoords{Tile: x + 1, Row: y})
		}

	case eChess.BISHOP:
		//towards bottom
		if rowLen < y {

		}

		//towards top
		if rowLen < 1 {

		}
	}
	return res, nil
}

func (g GameStateService) DrawCard(game *entities.Game, playerCode string) (*entities.Card, error) {
	return nil, errors.New("not implemented")
}

func (g GameStateService) CheckWin(game *entities.Game) (string, error) {
	return "", errors.New("not implemented")
}
