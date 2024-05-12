package board

import (
	"luckyChess/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _gameStoreService interfaces.IGameStoreService

func Register(router *gin.Engine, gameStoreService interfaces.IGameStoreService) {

	r := router.Group("/board")
	r.GET("", getBoard)

	_gameStoreService = gameStoreService
}

func getBoard(context *gin.Context) {

	game := _gameStoreService.GetGame("1")

	context.HTML(
		http.StatusOK,
		"board.html",
		game,
	)
}
