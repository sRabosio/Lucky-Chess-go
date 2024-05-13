package board

import (
	"luckyChess/entities"
	"luckyChess/services/interfaces"
	"net/http"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var _gameStoreService interfaces.IGameStoreService
var _gameTemplateService interfaces.IGameTemplatesService

func Register(router *gin.Engine, gameStoreService interfaces.IGameStoreService, gameTemplateService interfaces.IGameTemplatesService) {

	r := router.Group("/board")
	r.GET("", getBoard)

	_gameStoreService = gameStoreService
	_gameTemplateService = gameTemplateService
}

func getBoard(context *gin.Context) {

	store, err := cookie.Store.Get(cookie.NewStore(), context.Request, "gameCode")

	var game entities.Game

	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	val := store.Values["code"]

	if val == nil {
		game = _gameStoreService.NewGame(
			_gameTemplateService.GetTemplate("default"),
		)
	} else {
		game = _gameStoreService.GetGame(val.(string))
	}

	context.HTML(
		http.StatusOK,
		"board.html",
		game,
	)
}
