package gameStore

import (
	"luckyChess/entities"
	"strconv"
)

type GameStoreService struct {
	gameList map[string]entities.Game
}

func New() *GameStoreService {
	return &GameStoreService{
		gameList: make(map[string]entities.Game),
	}
}

func (g GameStoreService) NewGame(startingSet entities.BoardTemplate) entities.Game {

	board := entities.Board{
		Rows: [8]entities.Row{},
	}

	for plIndex, t := range startingSet.Template {
		for rowKey, row := range t {
			var currentRow *entities.Row = &board.Rows[rowKey]
			for tileKey, _ := range row {
				currentRow.Tiles[tileKey] = entities.Tile{
					Piece:    1,
					PlayerID: strconv.Itoa(plIndex),
				}
			}
		}
	}

	newGame := entities.Game{
		Board: board,
	}

	g.gameList["1"] = newGame
	return g.gameList["1"]
}

func (g GameStoreService) GetGame(gameCode string) entities.Game {
	return g.gameList[gameCode]
}

func (g GameStoreService) KillGame(gameCode string) bool {
	delete(g.gameList, gameCode)
	return true
}

func (g GameStoreService) ApplyChanges(code string, game entities.Game) bool {
	return false
}
