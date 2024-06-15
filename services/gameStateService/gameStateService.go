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
	if y > rowNum-1 {
		return res, errors.New("out of bounds")
	}

	tileNum := len(game.Board.Rows[y].Tiles)
	if x > tileNum-1 {
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

		res := []entities.TileCoords{}

		rows := game.Board.Rows

		//upward movement
		for i := y - 1; i > -1; i-- {
			if rows[i].Tiles[x].Piece > 0 {
				break
			}
			res = append(res, entities.TileCoords{Row: i, Tile: x})
		}

		//downward movement
		for i := y + 1; i < len(rows)-1; i++ {
			if rows[i].Tiles[x].Piece > 0 {
				break
			}
			res = append(res, entities.TileCoords{Row: i, Tile: x})
		}

		tiles := rows[y].Tiles

		//eastward movement
		for i := x + 1; i < len(tiles)-1; i++ {
			if tiles[i].Piece > 0 {
				break
			}
			res = append(res, entities.TileCoords{Row: y, Tile: i})
		}

		//westward movement
		for i := x - 1; i > -1; i-- {
			if tiles[i].Piece > 0 {
				break
			}
			res = append(res, entities.TileCoords{Row: y, Tile: i})
		}

		return res, nil
	},
}
