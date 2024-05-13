package gameStore

import (
	"luckyChess/entities"
	eChess "luckyChess/entities/EChess"
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

	//init template
	for plIndex, t := range startingSet.Template {
		for rowKey, row := range t {
			var currentRow *entities.Row = &board.Rows[rowKey]
			for tileKey, tile := range row {

				//TODO: use enum to convert string from json to int

				currentRow.Tiles[tileKey] = entities.Tile{
					Piece:    eChess.Parse(tile),
					PlayerID: strconv.Itoa(plIndex),
				}
			}
		}
	}

	//TODO: assign actual players

	newGame := entities.Game{
		Board: board,
	}

	g.gameList["1"] = newGame
	return newGame
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
