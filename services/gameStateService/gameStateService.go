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

func (g GameStateService) MovePiece(game *entities.Game, playerCode string, pieceCoords entities.TileCoords, targetCoords entities.TileCoords) error {
	moveset, err := g.GetMoveset(game, playerCode, pieceCoords)
	if err != nil {
		return err
	}

	isMoveValid := false
	for _, coords := range moveset {
		if targetCoords.Row == coords.Row && targetCoords.Tile == coords.Tile {
			isMoveValid = true
		}
	}
	if !isMoveValid {
		return errors.New("MovePieceService -> invalid move")
	}

	piece := game.Board.Rows[pieceCoords.Row].Tiles[pieceCoords.Tile].Piece
	plId := game.Board.Rows[pieceCoords.Row].Tiles[pieceCoords.Tile].PlayerID

	game.Board.Rows[pieceCoords.Row].Tiles[pieceCoords.Tile].Piece = 0
	game.Board.Rows[pieceCoords.Row].Tiles[pieceCoords.Tile].PlayerID = ""

	game.Board.Rows[targetCoords.Row].Tiles[targetCoords.Tile].Piece = piece
	game.Board.Rows[targetCoords.Row].Tiles[targetCoords.Tile].PlayerID = plId
	return nil
}

func (g GameStateService) GetMoveset(game *entities.Game, playerCode string, pieceCoords entities.TileCoords) ([]entities.TileCoords, error) {
	x, y := pieceCoords.Tile, pieceCoords.Row

	selectedTile := &game.Board.Rows[y].Tiles[x]

	res := []entities.TileCoords{}
	//TODO CHECK PLAYER TYPE: ex(first or second/ black or white)
	//or not actually: player's POV will always be the bottom pieces?????
	//board will need to be flipped at some point for someone

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

	//queen = bishop + rook
	//also can't reference map inside itself :(
	if selectedTile.Piece == eChess.QUEEN {
		res := []entities.TileCoords{}
		currRes, err := chessMoveset[eChess.BISHOP](game, x, y, playerCode)
		if err != nil {
			return nil, err
		}
		res = append(res, currRes...)

		currRes, err = chessMoveset[eChess.ROOK](game, x, y, playerCode)
		if err != nil {
			return nil, err
		}
		res = append(res, currRes...)

		return res, nil
	}

	movesetGetter := chessMoveset[selectedTile.Piece]
	if movesetGetter == nil {
		return nil, errors.New("missing getter for piece type " + selectedTile.Piece.String())
	}

	res, err := movesetGetter(game, x, y, playerCode)

	return res, err
}

func (g GameStateService) DrawCard(game *entities.Game, playerCode string) (*entities.Card, error) {
	return nil, errors.New("not implemented")
}

func (g GameStateService) CheckWin(game *entities.Game) (string, error) {
	return "", errors.New("not implemented")
}

func trySetCoords(game *entities.Game, xCurr int, yCurr int, plID string) (*entities.TileCoords, bool, error) {
	if yCurr < 0 || yCurr > len(game.Board.Rows)-1 {
		return nil, false, errors.New("invalid row")
	}

	if xCurr < 0 || xCurr > len(game.Board.Rows[yCurr].Tiles)-1 {
		return nil, false, errors.New("invalid tile")
	}

	if game.Board.Rows[yCurr].Tiles[xCurr].Piece > 0 {
		if game.Board.Rows[yCurr].Tiles[xCurr].PlayerID == plID {
			return nil, true, errors.New("player cannot eat themself")
		}
		return &entities.TileCoords{Row: yCurr, Tile: xCurr}, true, nil
	}

	return &entities.TileCoords{Row: yCurr, Tile: xCurr}, false, nil
}

type movesetGetter func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error)

// NB: coordinates are calculated from bottom player prespective
var chessMoveset = map[eChess.EChess]movesetGetter{
	eChess.PAWN: func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error) {
		res := []entities.TileCoords{}

		//in front
		coords, hasPiece, err := trySetCoords(game, x, y-1, plID)
		if err == nil && !hasPiece {
			res = append(res, *coords)
		}

		//right
		coords, hasPiece, err = trySetCoords(game, x-1, y-1, plID)
		if err == nil && hasPiece {
			res = append(res, *coords)
		}

		//left
		coords, hasPiece, err = trySetCoords(game, x+1, y-1, plID)
		if err == nil && hasPiece {
			res = append(res, *coords)
		}

		return res, nil
	},
	eChess.ROOK: func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error) {

		res := []entities.TileCoords{}

		rows := game.Board.Rows

		//upward movement
		for i := y - 1; i > -1; i-- {
			coords, hasPiece, err := trySetCoords(game, x, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		//downward movement
		for i := y + 1; i < len(rows)-1; i++ {
			coords, hasPiece, err := trySetCoords(game, x, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		tiles := rows[y].Tiles

		//eastward movement
		for i := x + 1; i < len(tiles)-1; i++ {
			coords, hasPiece, err := trySetCoords(game, i, y, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		//westward movement
		for i := x - 1; i > -1; i-- {
			coords, hasPiece, err := trySetCoords(game, i, y, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		return res, nil
	},
	eChess.KNIGHT: func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error) {
		res := []entities.TileCoords{}

		//up left
		coords, _, err := trySetCoords(game, x-2, y-1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		coords, _, err = trySetCoords(game, x-1, y-2, plID)
		if err == nil {
			res = append(res, *coords)
		}

		//up right
		coords, _, err = trySetCoords(game, x+2, y-1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		coords, _, err = trySetCoords(game, x+1, y-2, plID)
		if err == nil {
			res = append(res, *coords)
		}

		//down right
		coords, _, err = trySetCoords(game, x+2, y+1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		coords, _, err = trySetCoords(game, x+1, y+2, plID)
		if err == nil {
			res = append(res, *coords)
		}

		//down left
		coords, _, err = trySetCoords(game, x-2, y+1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		coords, _, err = trySetCoords(game, x-1, y+2, plID)
		if err == nil {
			res = append(res, *coords)
		}

		return res, nil
	},
	eChess.BISHOP: func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error) {

		res := []entities.TileCoords{}

		currentX := x
		//up left
		for i := y - 1; i > -1; i-- {
			currentX--
			coords, hasPiece, err := trySetCoords(game, currentX, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		currentX = x
		//up right
		for i := y - 1; i > -1; i-- {
			currentX++
			coords, hasPiece, err := trySetCoords(game, currentX, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		//down left
		currentX = x
		for i := y + 1; i < len(game.Board.Rows); i++ {
			currentX--
			coords, hasPiece, err := trySetCoords(game, currentX, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		//donw right
		currentX = x
		for i := y + 1; i < len(game.Board.Rows); i++ {
			currentX++
			coords, hasPiece, err := trySetCoords(game, currentX, i, plID)
			if err != nil {
				break
			}
			res = append(res, *coords)
			if hasPiece {
				break
			}
		}

		return res, nil
	},
	eChess.KING: func(game *entities.Game, x int, y int, plID string) ([]entities.TileCoords, error) {
		//todo: avoid movement where eaten

		res := []entities.TileCoords{}

		//left
		coords, _, err := trySetCoords(game, x-1, y, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//right
		coords, _, err = trySetCoords(game, x+1, y, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//topleft
		coords, _, err = trySetCoords(game, x-1, y-1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//topright
		coords, _, err = trySetCoords(game, x+1, y-1, plID)
		if err == nil {
			res = append(res, *coords)
		} //top
		coords, _, err = trySetCoords(game, x, y-1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//bottom
		coords, _, err = trySetCoords(game, x, y+1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//bottomleft
		coords, _, err = trySetCoords(game, x-1, y+1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		//bottomright
		coords, _, err = trySetCoords(game, x+1, y+1, plID)
		if err == nil {
			res = append(res, *coords)
		}
		return res, nil
	},
}
