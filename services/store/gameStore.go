package gameStore

import (
	"luckyChess/entities"
)

type GameStoreService struct {
	gameList map[string]entities.Game
}

func New() *GameStoreService {
	return &GameStoreService{
		gameList: make(map[string]entities.Game),
	}
}

func (g GameStoreService) NewGame(startingSet entities.Board) entities.Game {

	newGame := entities.Game{
		Board: startingSet,
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
