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

	rowNum := len(game.Board.Rows)
	if y >= rowNum-1 {
		return res, errors.New("out of bounds")
	}

	tileNum := len(game.Board.Rows[y].Tiles)
	if x >= tileNum-1 {
		return res, errors.New("out of bounds")
	}

	if selectedTile.Piece < 1 {
		return []entities.TileCoords{}, nil
	}

	if playerCode != selectedTile.PlayerID {
		return res, errors.New("invalid player")
	}

	movesetGetter := chessMoveset[selectedTile.Piece]
	if movesetGetter == nil {
		return nil, errors.New("missing getter for piece type " + selectedTile.Piece.String())
	}

	res, err := movesetGetter(game, x, y)

	return res, err
}

func (g GameStateService) DrawCard(game *entities.Game, playerCode string) (*entities.Card, error) {
	return nil, errors.New("not implemented")
}

func (g GameStateService) CheckWin(game *entities.Game) (string, error) {
	return "", errors.New("not implemented")
}

type movesetGetter func(game *entities.Game, x int, y int) ([]entities.TileCoords, error)

// NB: coordinates are calculated from bottom player prespective
var chessMoveset = map[eChess.EChess]movesetGetter{
	eChess.PAWN: func(game *entities.Game, x int, y int) ([]entities.TileCoords, error) {
		res := []entities.TileCoords{}

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

		return res, nil
	},
	eChess.ROOK: func(game *entities.Game, x int, y int) ([]entities.TileCoords, error) {
		// rows := slices.Clone()
		// slices.Reverse()
		// //column movement
		// for i, row := range game.Board.Rows {

		// }
		return []entities.TileCoords{}, nil
	},
}
