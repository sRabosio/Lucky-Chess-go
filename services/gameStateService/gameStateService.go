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
		currRes, err := chessMoveset[eChess.BISHOP](game, x, y)
		if err != nil {
			return nil, err
		}
		res = append(res, currRes...)

		currRes, err = chessMoveset[eChess.ROOK](game, x, y)
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

	res, err := movesetGetter(game, x, y)

	return res, err
}

func (g GameStateService) DrawCard(game *entities.Game, playerCode string) (*entities.Card, error) {
	return nil, errors.New("not implemented")
}

func (g GameStateService) CheckWin(game *entities.Game) (string, error) {
	return "", errors.New("not implemented")
}

func trySetCoords(game *entities.Game, xCurr int, yCurr int) (bool, *entities.TileCoords) {
	if yCurr < 0 || yCurr > len(game.Board.Rows)-1 {
		return false, nil
	}

	if xCurr < 0 || xCurr > len(game.Board.Rows[yCurr].Tiles)-1 {
		return false, nil
	}

	if game.Board.Rows[yCurr].Tiles[xCurr].Piece > 0 {
		return false, nil
	}
	return true, &entities.TileCoords{Row: yCurr, Tile: xCurr}
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
	eChess.KNIGHT: func(game *entities.Game, x int, y int) ([]entities.TileCoords, error) {
		res := []entities.TileCoords{}

		//up left
		valid, coords := trySetCoords(game, x-2, y-1)
		if valid {
			res = append(res, *coords)
		}
		valid, coords = trySetCoords(game, x-1, y-2)
		if valid {
			res = append(res, *coords)
		}

		//up right
		valid, coords = trySetCoords(game, x+2, y-1)
		if valid {
			res = append(res, *coords)
		}
		valid, coords = trySetCoords(game, x+1, y-2)
		if valid {
			res = append(res, *coords)
		}

		//down right
		valid, coords = trySetCoords(game, x+2, y+1)
		if valid {
			res = append(res, *coords)
		}
		valid, coords = trySetCoords(game, x+1, y+2)
		if valid {
			res = append(res, *coords)
		}

		//down left
		valid, coords = trySetCoords(game, x-2, y+1)
		if valid {
			res = append(res, *coords)
		}
		valid, coords = trySetCoords(game, x-1, y+2)
		if valid {
			res = append(res, *coords)
		}

		return res, nil
	},
	eChess.BISHOP: func(game *entities.Game, x int, y int) ([]entities.TileCoords, error) {

		res := []entities.TileCoords{}

		currentX := x
		//up left
		for i := y - 1; i > -1; i-- {
			currentX--
			valid, coords := trySetCoords(game, currentX, i)
			if !valid {
				break
			}
			res = append(res, *coords)
		}

		currentX = x
		//up right
		for i := y - 1; i > -1; i-- {
			currentX++
			valid, coords := trySetCoords(game, currentX, i)
			if !valid {
				break
			}
			res = append(res, *coords)
		}

		//down left
		currentX = x
		for i := y + 1; i < len(game.Board.Rows); i++ {
			currentX--
			valid, coords := trySetCoords(game, currentX, i)
			if !valid {
				break
			}
			res = append(res, *coords)
		}

		//donw right
		currentX = x
		for i := y + 1; i < len(game.Board.Rows); i++ {
			currentX++
			valid, coords := trySetCoords(game, currentX, i)
			if !valid {
				break
			}
			res = append(res, *coords)
		}

		return res, nil
	},
	eChess.KING: func(game *entities.Game, x int, y int) ([]entities.TileCoords, error) {
		//todo: avoid movement where eaten

		res := []entities.TileCoords{}

		//left
		valid, coords := trySetCoords(game, x-1, y)
		if valid {
			res = append(res, *coords)
		}
		//right
		valid, coords = trySetCoords(game, x+1, y)
		if valid {
			res = append(res, *coords)
		}
		//topleft
		valid, coords = trySetCoords(game, x-1, y-1)
		if valid {
			res = append(res, *coords)
		}
		//topright
		valid, coords = trySetCoords(game, x+1, y-1)
		if valid {
			res = append(res, *coords)
		} //top
		valid, coords = trySetCoords(game, x, y-1)
		if valid {
			res = append(res, *coords)
		}
		//bottom
		valid, coords = trySetCoords(game, x, y+1)
		if valid {
			res = append(res, *coords)
		}
		//bottomleft
		valid, coords = trySetCoords(game, x-1, y+1)
		if valid {
			res = append(res, *coords)
		}
		//bottomright
		valid, coords = trySetCoords(game, x+1, y+1)
		if valid {
			res = append(res, *coords)
		}
		return res, nil
	},
}
