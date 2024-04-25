package gameStore

import (
	"luckyChess/entities"
)

var gameList = make(map[string]entities.Game)

func NewGame(gameCode string, startingSet entities.Board) entities.Game {
	newGame := entities.Game{
		Board: startingSet,
	}

	gameList[gameCode] = newGame
	return gameList[gameCode]
}

func GetGame(gameCode string) entities.Game {
	return gameList[gameCode]
}

func KillGame(gameCode string) {
	delete(gameList, gameCode)
}
